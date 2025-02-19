package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type IVideoRepositoryWriter interface {
	CreateVideo(ctx context.Context, video *ent.Videos) (*ent.Videos, error)
	UpdateVideo(ctx context.Context, video *ent.Videos) (*ent.Videos, error)
	DeleteVideo(ctx context.Context, id uint) error
	CreateInBulk(ctx context.Context, videos []*ent.Videos) error
	GetMaxVersion(ctx context.Context) (int64, error)

	GetVideoById(ctx context.Context, id uint) (*ent.Videos, error)
}
