package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
)

type ICommentUsecase interface {
	GetCommentById(ctx context.Context, id uint) (*ent.Comments, error)
	GetListCommentByVideoID(ctx context.Context, videoID uint, key string) (*models.CommentsListResponse, error)
}
