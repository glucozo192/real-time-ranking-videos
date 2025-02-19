package usecase

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
)

type IViewerUsecase interface {
	GetViewerById(ctx context.Context, id uint) (*ent.Viewers, error)
	GetListViewerByVideoID(ctx context.Context, videoID uint, key string) (*models.ViewerListResponse, error)
}
