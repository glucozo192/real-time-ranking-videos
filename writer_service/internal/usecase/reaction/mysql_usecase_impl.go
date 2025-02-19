package usecase

import (
	"context"

	kafkaClient "github.com/glu/video-real-time-ranking/core/pkg/kafka"
	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/repositories"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/usecase"
	"github.com/pkg/errors"
)

type reactionUsecase struct {
	log                logger.Logger
	cfg                *config.Config
	reactionRepository repositories.IReactionRepositoryWriter
	kafkaProducer      kafkaClient.Producer
}

func NewReactionUsecase(log logger.Logger, cfg *config.Config, reactionRepository repositories.IReactionRepositoryWriter, kafkaProducer kafkaClient.Producer) usecase.IReactionUsecase {
	return &reactionUsecase{log: log, cfg: cfg, reactionRepository: reactionRepository, kafkaProducer: kafkaProducer}
}

func (p *reactionUsecase) CreateInBulk(ctx context.Context, reactions []*ent.Reactions) error {
	err := p.reactionRepository.CreateInBulk(ctx, reactions)
	if err != nil {
		return errors.Wrap(err, "reactionRepository.CreateInBulk")
	}
	return nil
}

func (p *reactionUsecase) Delete(ctx context.Context, id uint) error {
	err := p.reactionRepository.DeleteReaction(ctx, id)
	if err != nil {
		return errors.Wrap(err, "reactionRepository.DeleteReaction")
	}
	return nil
}

func (p *reactionUsecase) Update(ctx context.Context, reaction *ent.Reactions) (*ent.Reactions, error) {
	reactionDB, err := p.reactionRepository.UpdateReaction(ctx, reaction)
	if err != nil {
		return nil, errors.Wrap(err, "reactionRepository.UpdateReaction")
	}

	return reactionDB, nil
}

func (p *reactionUsecase) Create(ctx context.Context, reaction *ent.Reactions) (*ent.Reactions, error) {
	reactionDB, err := p.reactionRepository.CreateReaction(ctx, reaction)
	if err != nil {
		return nil, errors.Wrap(err, "reactionRepository.createdReaction")
	}

	return reactionDB, nil
}
