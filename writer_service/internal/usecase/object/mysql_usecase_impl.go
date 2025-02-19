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

type objectUsecase struct {
	log              logger.Logger
	cfg              *config.Config
	objectRepository repositories.IObjectRepositoryWriter
	kafkaProducer    kafkaClient.Producer
}

func NewObjectUsecase(log logger.Logger, cfg *config.Config, objectRepository repositories.IObjectRepositoryWriter, kafkaProducer kafkaClient.Producer) usecase.IObjectUsecase {
	return &objectUsecase{log: log, cfg: cfg, objectRepository: objectRepository, kafkaProducer: kafkaProducer}
}

func (p *objectUsecase) Delete(ctx context.Context, id uint) error {
	err := p.objectRepository.DeleteObject(ctx, id)
	if err != nil {
		return errors.Wrap(err, "objectRepository.DeleteObject")
	}
	return nil
}

func (p *objectUsecase) Update(ctx context.Context, object *ent.Objects) (*ent.Objects, error) {
	objectDB, err := p.objectRepository.UpdateObject(ctx, object)
	if err != nil {
		return nil, errors.Wrap(err, "objectRepository.UpdateObject")
	}

	return objectDB, nil
}

func (p *objectUsecase) Create(ctx context.Context, object *ent.Objects) (*ent.Objects, error) {
	objectDB, err := p.objectRepository.CreateObject(ctx, object)
	if err != nil {
		return nil, errors.Wrap(err, "objectRepository.createdObject")
	}

	return objectDB, nil
}
