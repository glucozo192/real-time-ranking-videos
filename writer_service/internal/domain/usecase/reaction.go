package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type IReactionUsecase interface {
	Create(ctx context.Context, Reaction *ent.Reactions) (*ent.Reactions, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, Reaction *ent.Reactions) (*ent.Reactions, error)
	CreateInBulk(ctx context.Context, reactions []*ent.Reactions) error

	//GetByID(ctx context.Context, query *models.GetProductByIdQuery) (*models.Product, error)
}
