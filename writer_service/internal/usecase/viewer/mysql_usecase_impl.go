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

type viewerUsecase struct {
	log              logger.Logger
	cfg              *config.Config
	viewerRepository repositories.IViewerRepositoryWriter
	kafkaProducer    kafkaClient.Producer
}

func NewViewerUsecase(log logger.Logger, cfg *config.Config, viewerRepository repositories.IViewerRepositoryWriter, kafkaProducer kafkaClient.Producer) usecase.IViewerUsecase {
	return &viewerUsecase{log: log, cfg: cfg, viewerRepository: viewerRepository, kafkaProducer: kafkaProducer}
}

func (p *viewerUsecase) CreateInBulk(ctx context.Context, viewers []*ent.Viewers) error {
	err := p.viewerRepository.CreateInBulk(ctx, viewers)
	if err != nil {
		return errors.Wrap(err, "viewerRepository.CreateInBulk")
	}
	return nil
}

func (p *viewerUsecase) Delete(ctx context.Context, id uint) error {
	err := p.viewerRepository.DeleteViewer(ctx, id)
	if err != nil {
		return errors.Wrap(err, "viewerRepository.DeleteViewer")
	}
	return nil
}

func (p *viewerUsecase) Update(ctx context.Context, viewer *ent.Viewers) (*ent.Viewers, error) {
	viewerDB, err := p.viewerRepository.UpdateViewer(ctx, viewer)
	if err != nil {
		return nil, errors.Wrap(err, "viewerRepository.UpdateViewer")
	}

	return viewerDB, nil
}

func (p *viewerUsecase) Create(ctx context.Context, viewer *ent.Viewers) (*ent.Viewers, error) {
	viewerDB, err := p.viewerRepository.CreateViewer(ctx, viewer)
	if err != nil {
		return nil, errors.Wrap(err, "viewerRepository.createdViewer")
	}

	return viewerDB, nil
}
