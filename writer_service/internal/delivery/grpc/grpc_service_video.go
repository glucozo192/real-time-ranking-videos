package grpc

import (
	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	product2 "github.com/glu/video-real-time-ranking/writer_service/internal/metrics"
	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcService struct {
	log     logger.Logger
	cfg     *config.Config
	v       *validator.Validate
	metrics *product2.WriterServiceMetrics
}

func NewWriterGrpcService(
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	metrics *product2.WriterServiceMetrics,
) *grpcService {
	return &grpcService{
		log:     log,
		cfg:     cfg,
		v:       v,
		metrics: metrics,
	}
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	s.metrics.ErrorGrpcRequests.Inc()
	return status.Error(c, err.Error())
}
