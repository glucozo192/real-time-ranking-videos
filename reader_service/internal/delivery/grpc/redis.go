package grpc

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
)

func (s *grpcService) RemoveCachingByKey(ctx context.Context, req *readerService.RemoveCachingByKeyReq) (*readerService.RemoveCachingByKeyRes, error) {
	s.metrics.CreateProductGrpcRequests.Inc()

	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "grpcService.CreateProduct")
	defer span.Finish()

	if err := s.redisClient.HDel(ctx, req.PrefixKey, req.Key).Err(); err != nil {
		s.log.WarnMsg("redisClient.HDel", err)
		return nil, err
	}
	s.log.Debugf("HDel prefix: %s, key: %s", req.PrefixKey, req.Key)

	s.metrics.SuccessGrpcRequests.Inc()
	return &readerService.RemoveCachingByKeyRes{}, nil
}
