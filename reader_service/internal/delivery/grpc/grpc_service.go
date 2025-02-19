package grpc

import (
	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/metrics"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcService struct {
	log         logger.Logger
	cfg         *config.Config
	v           *validator.Validate
	redisClient redis.UniversalClient
	metrics     *metrics.ReaderServiceMetrics
}

func NewReaderGrpcService(
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	redisClient redis.UniversalClient,
	metrics *metrics.ReaderServiceMetrics,
) *grpcService {
	return &grpcService{
		log:         log,
		cfg:         cfg,
		v:           v,
		redisClient: redisClient,
		metrics:     metrics,
	}
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	s.metrics.ErrorGrpcRequests.Inc()
	return status.Error(c, err.Error())
}
