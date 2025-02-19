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

type reactionUsecase struct {
	log       logger.Logger
	cfg       *config.Config
	redisRepo repositories.IReactionCacheRepository
	entRepo   repositories.IReactionRepositoryReader
}

func NewReactionUsecase(
	log logger.Logger,
	cfg *config.Config,
	redisRepo repositories.IReactionCacheRepository,
	entRepo repositories.IReactionRepositoryReader,
) usecase.IReactionUsecase {
	return &reactionUsecase{
		log:       log,
		cfg:       cfg,
		redisRepo: redisRepo,
		entRepo:   entRepo,
	}
}

func (v *reactionUsecase) GetReactionById(ctx context.Context, id uint) (*ent.Reactions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getReactionByIdHandler.Handle")
	defer span.Finish()

	if reaction, err := v.redisRepo.GetReaction(ctx, strconv.Itoa(int(id))); err == nil && reaction != nil {
		return reaction, nil
	}

	reactionResp, err := v.entRepo.GetReactionById(ctx, id)
	if err != nil {
		return nil, err
	}

	v.redisRepo.PutReaction(ctx, strconv.Itoa(int(id)), reactionResp)
	return reactionResp, nil
}

func (v *reactionUsecase) GetListReactionByVideoID(ctx context.Context, videoID uint, key string) (*models.ReactionsListResponse, error) {
	//if reactionCacheResp, err := v.redisRepo.GetReactionsByKey(ctx, key); err == nil && reactionCacheResp.Reactions != nil {
	//	return &reactionCacheResp, nil
	//}

	reactionResp, err := v.entRepo.GetListReactionByVideoID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	//if err = v.redisRepo.PutReactions(ctx, key, *reactionResp); err != nil {
	//	return nil, err
	//}
	return reactionResp, nil
}
