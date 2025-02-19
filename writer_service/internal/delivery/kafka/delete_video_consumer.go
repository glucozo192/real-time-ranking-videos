package kafka

import (
	"context"

	"github.com/avast/retry-go"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
	"github.com/glu/video-real-time-ranking/core/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *MessageProcessor) processDeleteVideo(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.DeleteVideoKafkaMessages.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "messageProcessor.processDeleteVideo")
	defer span.Finish()

	msg := &kafkaMessages.VideoDelete{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		return s.videoUsecase.Delete(ctx, uint(msg.GetVideoID()))
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("DeleteVideo.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}

	s.commitMessage(ctx, r, m)
}
