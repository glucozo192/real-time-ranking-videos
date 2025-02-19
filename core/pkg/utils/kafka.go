package utils

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"

	kafkaClient "github.com/glu/video-real-time-ranking/core/pkg/kafka"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"

	"google.golang.org/protobuf/proto"
)

func PublishKafkaMessage(ctx context.Context, span opentracing.SpanContext, kafkaProducer kafkaClient.Producer, msg proto.Message, topicName string) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Topic:   topicName,
		Value:   msgBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span),
	}

	return kafkaProducer.PublishMessage(ctx, message)
}
