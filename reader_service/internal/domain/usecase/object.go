package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"

	"github.com/glu/video-real-time-ranking/ent"
)

type IObjectUsecase interface {
	GetObjectById(ctx context.Context, id uint) (*ent.Objects, error)
	GetListObjectByVideoID(ctx context.Context, videoID uint, key string) (*models.ObjectsListResponse, error)
	GetListObjectByVideoPath(ctx context.Context, path string) ([]*ent.Objects, error)
}
