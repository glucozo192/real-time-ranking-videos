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

type commentUsecase struct {
	log               logger.Logger
	cfg               *config.Config
	commentRepository repositories.ICommentRepositoryWriter
	kafkaProducer     kafkaClient.Producer
}

func NewCommentUsecase(log logger.Logger, cfg *config.Config, commentRepository repositories.ICommentRepositoryWriter, kafkaProducer kafkaClient.Producer) usecase.ICommentUsecase {
	return &commentUsecase{log: log, cfg: cfg, commentRepository: commentRepository, kafkaProducer: kafkaProducer}
}

func (p *commentUsecase) CreateInBulk(ctx context.Context, comments []*ent.Comments) error {
	err := p.commentRepository.CreateInBulk(ctx, comments)
	if err != nil {
		return errors.Wrap(err, "commentRepository.CreateInBulk")
	}
	return nil
}

func (p *commentUsecase) Delete(ctx context.Context, id uint) error {
	err := p.commentRepository.DeleteComment(ctx, id)
	if err != nil {
		return errors.Wrap(err, "commentRepository.DeleteComment")
	}
	return nil
}

func (p *commentUsecase) Update(ctx context.Context, comment *ent.Comments) (*ent.Comments, error) {
	commentDB, err := p.commentRepository.UpdateComment(ctx, comment)
	if err != nil {
		return nil, errors.Wrap(err, "commentRepository.UpdateComment")
	}

	return commentDB, nil
}

func (p *commentUsecase) Create(ctx context.Context, comment *ent.Comments) (*ent.Comments, error) {
	commentDB, err := p.commentRepository.CreateComment(ctx, comment)
	if err != nil {
		return nil, errors.Wrap(err, "commentRepository.createdComment")
	}

	return commentDB, nil
}
