package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
)

type IObjectRepositoryWriter interface {
	CreateObject(ctx context.Context, Object *ent.Objects) (*ent.Objects, error)
	UpdateObject(ctx context.Context, Object *ent.Objects) (*ent.Objects, error)
	DeleteObject(ctx context.Context, id uint) error

	GetListObjectByVideoID(ctx context.Context, videoID uint) ([]*ent.Objects, error)
	CreateInBulk(ctx context.Context, objects []*ent.Objects) error
}
