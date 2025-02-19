package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
)

type ICommentRepositoryReader interface {
	GetCommentById(ctx context.Context, id uint) (*ent.Comments, error)
	GetListCommentByVideoID(ctx context.Context, videoID uint) (*models.CommentsListResponse, error)
}

type ICommentCacheRepository interface {
	PutComment(ctx context.Context, key string, comment *ent.Comments)
	GetComment(ctx context.Context, key string) (*ent.Comments, error)
	DelComment(ctx context.Context, key string)
	PutComments(ctx context.Context, key string, comments models.CommentsListResponse) error
	GetCommentsByKey(ctx context.Context, key string) (models.CommentsListResponse, error)
	DelAllComments(ctx context.Context)
}
