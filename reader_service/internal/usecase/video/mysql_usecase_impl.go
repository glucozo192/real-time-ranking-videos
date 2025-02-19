package product

import (
	"context"
	"strconv"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	models2 "github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/usecase"
	"github.com/opentracing/opentracing-go"
)

type videoUsecase struct {
	log       logger.Logger
	cfg       *config.Config
	redisRepo repositories.IVideoCacheRepository
	entRepo   repositories.IVideoRepositoryReader
}

func NewVideoUsecase(
	log logger.Logger,
	cfg *config.Config,
	redisRepo repositories.IVideoCacheRepository,
	entRepo repositories.IVideoRepositoryReader,
) usecase.IVideoUsecase {
	return &videoUsecase{
		log:       log,
		cfg:       cfg,
		redisRepo: redisRepo,
		entRepo:   entRepo,
	}
}

func (v *videoUsecase) GetVideoById(ctx context.Context, id uint) (*ent.Videos, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getVideoByIdHandler.Handle")
	defer span.Finish()

	if video, err := v.redisRepo.GetVideo(ctx, strconv.Itoa(int(id))); err == nil && video != nil {
		return video, nil
	}

	videoResp, err := v.entRepo.GetVideoById(ctx, id)
	if err != nil {
		return nil, err
	}

	v.redisRepo.PutVideo(ctx, strconv.Itoa(int(id)), videoResp)
	return videoResp, nil
}

func (v *videoUsecase) SearchVideo(ctx context.Context, query models2.SearchVideoRequest, key string) (*models2.VideosListResponse, error) {
	//if videoCacheResp, err := v.redisRepo.GetVideosByKey(ctx, key); err == nil && videoCacheResp.Videos != nil {
	//	return &videoCacheResp, nil
	//}

	videoResp, err := v.entRepo.SearchVideoByParams(ctx, query)
	if err != nil {
		return nil, err
	}

	//if err = v.redisRepo.PutVideos(ctx, key, *videoResp); err != nil {
	//	return nil, err
	//}
	return videoResp, nil
}
