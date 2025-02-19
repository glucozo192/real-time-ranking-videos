package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
)

type IObjectRepositoryReader interface {
	GetObjectById(ctx context.Context, id uint) (*ent.Objects, error)
	GetListObjectByVideoID(ctx context.Context, videoID uint) (*models.ObjectsListResponse, error)
	GetListObjectByVideoIDV2(ctx context.Context, videoID uint) ([]*ent.Objects, error)
}

type IObjectCacheRepository interface {
	PutObject(ctx context.Context, key string, object *ent.Objects)
	GetObject(ctx context.Context, key string) (*ent.Objects, error)
	DelObject(ctx context.Context, key string)
	PutObjects(ctx context.Context, key string, objects models.ObjectsListResponse) error
	GetObjectsByKey(ctx context.Context, key string) (models.ObjectsListResponse, error)
	DelAllObjects(ctx context.Context)
}
