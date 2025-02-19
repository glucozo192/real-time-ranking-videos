package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"

	"github.com/glu/video-real-time-ranking/ent"
)

type IReactionUsecase interface {
	GetReactionById(ctx context.Context, id uint) (*ent.Reactions, error)
	GetListReactionByVideoID(ctx context.Context, videoID uint, key string) (*models.ReactionsListResponse, error)
}
