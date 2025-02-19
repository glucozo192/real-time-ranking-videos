package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type IViewerRepositoryWriter interface {
	CreateViewer(ctx context.Context, Viewer *ent.Viewers) (*ent.Viewers, error)
	UpdateViewer(ctx context.Context, Viewer *ent.Viewers) (*ent.Viewers, error)
	DeleteViewer(ctx context.Context, id uint) error
	CreateInBulk(ctx context.Context, viewers []*ent.Viewers) error

	GetListViewerByVideoID(ctx context.Context, videoID uint) ([]*ent.Viewers, error)
}
