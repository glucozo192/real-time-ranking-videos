package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type IReactionRepositoryWriter interface {
	CreateReaction(ctx context.Context, Reaction *ent.Reactions) (*ent.Reactions, error)
	UpdateReaction(ctx context.Context, Reaction *ent.Reactions) (*ent.Reactions, error)
	DeleteReaction(ctx context.Context, id uint) error
	CreateInBulk(ctx context.Context, reactions []*ent.Reactions) error

	GetListReactionByVideoID(ctx context.Context, videoID uint) ([]*ent.Reactions, error)
}
