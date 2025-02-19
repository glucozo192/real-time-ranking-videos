package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
)

type IVideoRepositoryReader interface {
	GetVideoById(ctx context.Context, id uint) (*ent.Videos, error)
	GetVideoByVideoUrl(ctx context.Context, url string) (*ent.Videos, error)
	SearchVideoByParams(ctx context.Context, search models.SearchVideoRequest) (*models.VideosListResponse, error)
}

type IVideoCacheRepository interface {
	PutVideo(ctx context.Context, key string, video *ent.Videos)
	GetVideo(ctx context.Context, key string) (*ent.Videos, error)
	DelVideo(ctx context.Context, key string)
	PutVideos(ctx context.Context, key string, videos models.VideosListResponse) error
	GetVideosByKey(ctx context.Context, key string) (models.VideosListResponse, error)
	DelAllVideos(ctx context.Context)
}
