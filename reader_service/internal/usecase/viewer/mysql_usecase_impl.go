package product

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/usecase"

	"strconv"

	"github.com/opentracing/opentracing-go"
)

type viewerUsecase struct {
	log       logger.Logger
	cfg       *config.Config
	redisRepo repositories.IViewerCacheRepository
	entRepo   repositories.IViewerRepositoryReader
}

func NewViewerUsecase(
	log logger.Logger,
	cfg *config.Config,
	redisRepo repositories.IViewerCacheRepository,
	entRepo repositories.IViewerRepositoryReader,
) usecase.IViewerUsecase {
	return &viewerUsecase{
		log:       log,
		cfg:       cfg,
		redisRepo: redisRepo,
		entRepo:   entRepo,
	}
}

func (v *viewerUsecase) GetViewerById(ctx context.Context, id uint) (*ent.Viewers, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getViewerByIdHandler.Handle")
	defer span.Finish()

	if viewer, err := v.redisRepo.GetViewer(ctx, strconv.Itoa(int(id))); err == nil && viewer != nil {
		return viewer, nil
	}

	viewerResp, err := v.entRepo.GetViewerById(ctx, id)
	if err != nil {
		return nil, err
	}

	v.redisRepo.PutViewer(ctx, strconv.Itoa(int(id)), viewerResp)
	return viewerResp, nil
}

func (v *viewerUsecase) GetListViewerByVideoID(ctx context.Context, videoID uint, key string) (*models.ViewerListResponse, error) {
	//if viewerCacheResp, err := v.redisRepo.GetViewersByKey(ctx, key); err == nil && viewerCacheResp.Viewers != nil {
	//	return &viewerCacheResp, nil
	//}

	viewerResp, err := v.entRepo.GetListViewerByVideoID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	//if err = v.redisRepo.PutViewers(ctx, key, *viewerResp); err != nil {
	//	return nil, err
	//}
	return viewerResp, nil
}
