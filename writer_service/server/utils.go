package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/glu/video-real-time-ranking/core/pkg/constants"
	kafkaClient "github.com/glu/video-real-time-ranking/core/pkg/kafka"
	"github.com/heptiolabs/healthcheck"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	stackSize = 1 << 10 // 1 KB
)

func (s *server) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, s.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaCon")
	}

	s.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	s.log.Infof("kafka connected to brokers: %+v", brokers)

	return nil
}

func (s *server) initKafkaTopics(ctx context.Context) {
	controller, err := s.kafkaConn.Controller()
	if err != nil {
		s.log.WarnMsg("kafkaConn.Controller", err)
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	s.log.Infof("kafka controller uri: %s", controllerURI)

	conn, err := kafka.DialContext(ctx, "tcp", controllerURI)
	if err != nil {
		s.log.WarnMsg("initKafkaTopics.DialContext", err)
		return
	}
	defer conn.Close() // nolint: errcheck

	s.log.Infof("established new kafka controller connection: %s", controllerURI)

	// init topic for video
	videoCreateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.VideoCreate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.VideoCreate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.VideoCreate.ReplicationFactor,
	}

	videoCreatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.VideoCreated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.VideoCreated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.VideoCreated.ReplicationFactor,
	}

	videoUpdateTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.VideoUpdate.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.VideoUpdate.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.VideoUpdate.ReplicationFactor,
	}

	videoUpdatedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.VideoUpdated.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.VideoUpdated.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.VideoUpdated.ReplicationFactor,
	}

	videoDeleteTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.VideoDelete.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.VideoDelete.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.VideoDelete.ReplicationFactor,
	}

	videoDeletedTopic := kafka.TopicConfig{
		Topic:             s.cfg.KafkaTopics.VideoDeleted.TopicName,
		NumPartitions:     s.cfg.KafkaTopics.VideoDeleted.Partitions,
		ReplicationFactor: s.cfg.KafkaTopics.VideoDeleted.ReplicationFactor,
	}

	if err := conn.CreateTopics(
		videoCreateTopic,
		videoUpdateTopic,
		videoCreatedTopic,
		videoUpdatedTopic,
		videoDeleteTopic,
		videoDeletedTopic,
	); err != nil {
		s.log.WarnMsg("kafkaConn.CreateTopics", err)
		return
	}

	s.log.Infof("kafka topics created or already exists: %+v",
		[]kafka.TopicConfig{
			videoCreateTopic,
			videoUpdateTopic,
			videoCreatedTopic,
			videoUpdatedTopic,
			videoDeleteTopic,
			videoDeletedTopic,
		})
}

func (s *server) getConsumerGroupTopics() []string {
	return []string{
		s.cfg.KafkaTopics.VideoCreate.TopicName,
		s.cfg.KafkaTopics.VideoUpdate.TopicName,
		s.cfg.KafkaTopics.VideoDelete.TopicName,
	}
}

func (s *server) runHealthCheck(ctx context.Context) {
	health := healthcheck.NewHandler()

	health.AddLivenessCheck(s.cfg.ServiceName, healthcheck.AsyncWithContext(ctx, func() error {
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	health.AddReadinessCheck(constants.Mysql, healthcheck.AsyncWithContext(ctx, func() error {
		// Assuming you have an "entClient" instance that connects to MySQL
		err := s.entClient.Schema.Create(ctx) // This can be any database-related operation
		if err != nil {
			return err
		}
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	health.AddReadinessCheck(constants.Kafka, healthcheck.AsyncWithContext(ctx, func() error {
		_, err := s.kafkaConn.Brokers()
		if err != nil {
			return err
		}
		return nil
	}, time.Duration(s.cfg.Probes.CheckIntervalSeconds)*time.Second))

	go func() {
		s.log.Infof("Writer microservice Kubernetes probes listening on port: %s", s.cfg.Probes.Port)
		if err := http.ListenAndServe(s.cfg.Probes.Port, health); err != nil {
			s.log.WarnMsg("ListenAndServe", err)
		}
	}()
}

func (s *server) runMetrics(cancel context.CancelFunc) {
	metricsServer := &fasthttp.Server{
		Handler: s.metricsHandler,
	}

	go func() {
		addr := fmt.Sprintf(":%s", s.cfg.Probes.PrometheusPort)
		if err := metricsServer.ListenAndServe(addr); err != nil {
			s.log.Printf("metricsServer.ListenAndServe: %v", err)
			cancel()
		}
	}()

	s.log.Printf("Metrics server is running on port: %s", s.cfg.Probes.PrometheusPort)
}

func (s *server) metricsHandler(ctx *fasthttp.RequestCtx) {
	// Your middleware/recovery logic here
	// Note: FastHTTP doesn't have a built-in middleware system like Echo
	// You need to implement the middleware logic manually if needed

	// Serve Prometheus metrics using the fasthttpadaptor and promhttp packages
	if string(ctx.Path()) == s.cfg.Probes.PrometheusPath {
		fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	}
}
