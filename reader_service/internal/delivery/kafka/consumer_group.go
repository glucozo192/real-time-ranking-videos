package kafka

import (
	"context"
	"sync"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/usecase"
	"github.com/glu/video-real-time-ranking/reader_service/internal/metrics"
	"github.com/go-playground/validator"
	"github.com/segmentio/kafka-go"
)

const (
	PoolSize = 30
)

type readerMessageProcessor struct {
	log          logger.Logger
	cfg          *config.Config
	v            *validator.Validate
	videoUsecase usecase.IVideoUsecase
	metrics      *metrics.ReaderServiceMetrics
}

func NewReaderMessageProcessor(
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	videoUsecase usecase.IVideoUsecase,
	metrics *metrics.ReaderServiceMetrics,
) *readerMessageProcessor {
	return &readerMessageProcessor{
		log:          log,
		cfg:          cfg,
		v:            v,
		videoUsecase: videoUsecase,
		metrics:      metrics,
	}
}

func (s *readerMessageProcessor) ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
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

		}
	}
}
