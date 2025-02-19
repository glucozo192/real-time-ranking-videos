package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/avast/retry-go"
	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/usecase"
	product2 "github.com/glu/video-real-time-ranking/writer_service/internal/metrics"
	"github.com/go-playground/validator"
	"github.com/segmentio/kafka-go"
)

const (
	PoolSize = 30
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

type MessageProcessor struct {
	log          logger.Logger
	cfg          *config.Config
	v            *validator.Validate
	videoUsecase usecase.IVideoUsecase
	metrics      *product2.WriterServiceMetrics
}

func NewMessageProcessor(
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	videoUsecase usecase.IVideoUsecase,
	metrics *product2.WriterServiceMetrics,
) *MessageProcessor {
	return &MessageProcessor{
		log:          log,
		cfg:          cfg,
		v:            v,
		videoUsecase: videoUsecase,
		metrics:      metrics,
	}
}

func (s *MessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}

		s.logProcessMessage(m, workerID)

		switch m.Topic {
		case s.cfg.KafkaTopics.VideoCreate.TopicName:
			s.processCreateVideo(ctx, r, m)
		case s.cfg.KafkaTopics.VideoUpdate.TopicName:
			s.processUpdateVideo(ctx, r, m)
		case s.cfg.KafkaTopics.VideoDelete.TopicName:
			s.processDeleteVideo(ctx, r, m)
		}
	}
}
