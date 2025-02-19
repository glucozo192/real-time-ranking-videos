package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type IViewerUsecase interface {
	Create(ctx context.Context, Viewer *ent.Viewers) (*ent.Viewers, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, Viewer *ent.Viewers) (*ent.Viewers, error)
	CreateInBulk(ctx context.Context, viewers []*ent.Viewers) error

	//GetByID(ctx context.Context, query *models.GetProductByIdQuery) (*models.Product, error)
}
