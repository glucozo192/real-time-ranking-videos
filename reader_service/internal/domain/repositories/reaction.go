package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
)

type IReactionRepositoryReader interface {
	GetReactionById(ctx context.Context, id uint) (*ent.Reactions, error)
	GetListReactionByVideoID(ctx context.Context, videoID uint) (*models.ReactionsListResponse, error)
}

type IReactionCacheRepository interface {
	PutReaction(ctx context.Context, key string, reaction *ent.Reactions)
	GetReaction(ctx context.Context, key string) (*ent.Reactions, error)
	DelReaction(ctx context.Context, key string)
	PutReactions(ctx context.Context, key string, reactions models.ReactionsListResponse) error
	GetReactionsByKey(ctx context.Context, key string) (models.ReactionsListResponse, error)
	DelAllReactions(ctx context.Context)
}
