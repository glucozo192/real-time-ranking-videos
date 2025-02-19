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

type objectUsecase struct {
	log       logger.Logger
	cfg       *config.Config
	redisRepo repositories.IObjectCacheRepository
	entRepo   repositories.IObjectRepositoryReader
	videoRepo repositories.IVideoRepositoryReader
}

func NewObjectUsecase(
	log logger.Logger,
	cfg *config.Config,
	redisRepo repositories.IObjectCacheRepository,
	entRepo repositories.IObjectRepositoryReader,
	videoRepo repositories.IVideoRepositoryReader,
) usecase.IObjectUsecase {
	return &objectUsecase{
		log:       log,
		cfg:       cfg,
		redisRepo: redisRepo,
		entRepo:   entRepo,
		videoRepo: videoRepo,
	}
}

func (v *objectUsecase) GetObjectById(ctx context.Context, id uint) (*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getObjectByIdHandler.Handle")
	defer span.Finish()

	if object, err := v.redisRepo.GetObject(ctx, strconv.Itoa(int(id))); err == nil && object != nil {
		return object, nil
	}

	objectResp, err := v.entRepo.GetObjectById(ctx, id)
	if err != nil {
		return nil, err
	}

	v.redisRepo.PutObject(ctx, strconv.Itoa(int(id)), objectResp)
	return objectResp, nil
}

func (v *objectUsecase) GetListObjectByVideoID(ctx context.Context, videoID uint, key string) (*models.ObjectsListResponse, error) {
	//if objectCacheResp, err := v.redisRepo.GetObjectsByKey(ctx, key); err == nil && objectCacheResp.Objects != nil {
	//	return &objectCacheResp, nil
	//}

	objectResp, err := v.entRepo.GetListObjectByVideoID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	//if err = v.redisRepo.PutObjects(ctx, key, *objectResp); err != nil {
	//	return nil, err
	//}
	return objectResp, nil
}

func (v *objectUsecase) GetListObjectByVideoPath(ctx context.Context, path string) ([]*ent.Objects, error) {
	video, err := v.videoRepo.GetVideoByVideoUrl(ctx, path)
	if err != nil {
		return nil, err
	}
	objectResp, err := v.entRepo.GetListObjectByVideoIDV2(ctx, video.ID)
	if err != nil {
		return nil, err
	}

	return objectResp, nil
}
