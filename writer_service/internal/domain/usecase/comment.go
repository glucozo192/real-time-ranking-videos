package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type ICommentUsecase interface {
	Create(ctx context.Context, comment *ent.Comments) (*ent.Comments, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, comment *ent.Comments) (*ent.Comments, error)
	CreateInBulk(ctx context.Context, comments []*ent.Comments) error

	//GetByID(ctx context.Context, query *models.GetProductByIdQuery) (*models.Product, error)
}
