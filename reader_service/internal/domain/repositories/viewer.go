package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
)

type IViewerRepositoryReader interface {
	GetViewerById(ctx context.Context, id uint) (*ent.Viewers, error)
	GetListViewerByVideoID(ctx context.Context, videoID uint) (*models.ViewerListResponse, error)
}

type IViewerCacheRepository interface {
	PutViewer(ctx context.Context, key string, viewer *ent.Viewers)
	GetViewer(ctx context.Context, key string) (*ent.Viewers, error)
	DelViewer(ctx context.Context, key string)
	PutViewers(ctx context.Context, key string, viewers models.ViewerListResponse) error
	GetViewersByKey(ctx context.Context, key string) (models.ViewerListResponse, error)
	DelAllViewers(ctx context.Context)
}
