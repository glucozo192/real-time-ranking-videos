package kafka

import (
	"context"

	"github.com/avast/retry-go"
	"github.com/glu/video-real-time-ranking/core/pkg/tracing"
	"github.com/glu/video-real-time-ranking/core/proto/kafka"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *MessageProcessor) processUpdateVideo(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metrics.UpdateVideoKafkaMessages.Inc()

	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "messageProcessor.processUpdateVideo")
	defer span.Finish()

	var msg kafkaMessages.VideoUpdate
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		_, err := s.videoUsecase.Update(ctx, &ent.Videos{
			ID:           uint(msg.GetVideoID()),
			Name:         msg.GetName(),
			Description:  msg.GetDescription(),
			VideoURL:     msg.GetVideoUrl(),
			Config:       msg.GetConfig(),
			PathResource: msg.GetPathResource(),
			LevelSystem:  msg.GetLevelSystem(),
			Status:       msg.GetStatus(),
			Note:         msg.GetNote(),
			Assign:       msg.GetAssign(),
			Author:       msg.GetAuthor(),
		})
		return err
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		s.log.WarnMsg("UpdateVideo.Handle", err)
		s.metrics.ErrorKafkaMessages.Inc()
		return
	}

	s.commitMessage(ctx, r, m)
}
