package usecase

import (
	"context"
	"mime/multipart"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/models"
)

type IVideoUsecase interface {
	Create(ctx context.Context, video *ent.Videos) (*ent.Videos, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, video *ent.Videos) (*ent.Videos, error)
	CreateInBulk(ctx context.Context, videos []*ent.Videos) error
	GetMaxVersion(ctx context.Context) (int64, error)

	//GetByID(ctx context.Context, query *models.GetProductByIdQuery) (*models.Product, error)
	SaveJsonVideoItem(ctx context.Context, video *ent.Videos, comments []*ent.Comments, viewers []*ent.Viewers, reactions []*ent.Reactions, objects []*ent.Objects) error
	CreateVideoItemConfig(ctx context.Context, videoId uint) error
	ZipById(ctx context.Context, videoId uint, resource *models.VideoItem, versionNew int) error
	BulkImportByExcel(ctx context.Context, file *multipart.FileHeader) error
}
