package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type IObjectUsecase interface {
	Create(ctx context.Context, Object *ent.Objects) (*ent.Objects, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, Object *ent.Objects) (*ent.Objects, error)

	//GetByID(ctx context.Context, query *models.GetProductByIdQuery) (*models.Product, error)
}
