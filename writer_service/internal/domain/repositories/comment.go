package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type ICommentRepositoryWriter interface {
	CreateComment(ctx context.Context, comment *ent.Comments) (*ent.Comments, error)
	UpdateComment(ctx context.Context, comment *ent.Comments) (*ent.Comments, error)
	DeleteComment(ctx context.Context, id uint) error
	CreateInBulk(ctx context.Context, comments []*ent.Comments) error

	GetListCommentByVideoID(ctx context.Context, videoID uint) ([]*ent.Comments, error)
}
