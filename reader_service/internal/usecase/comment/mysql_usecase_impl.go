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

type commentUsecase struct {
	log       logger.Logger
	cfg       *config.Config
	redisRepo repositories.ICommentCacheRepository
	entRepo   repositories.ICommentRepositoryReader
}

func NewCommentUsecase(
	log logger.Logger,
	cfg *config.Config,
	redisRepo repositories.ICommentCacheRepository,
	entRepo repositories.ICommentRepositoryReader,
) usecase.ICommentUsecase {
	return &commentUsecase{
		log:       log,
		cfg:       cfg,
		redisRepo: redisRepo,
		entRepo:   entRepo,
	}
}

func (v *commentUsecase) GetCommentById(ctx context.Context, id uint) (*ent.Comments, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getCommentByIdHandler.Handle")
	defer span.Finish()

	if comment, err := v.redisRepo.GetComment(ctx, strconv.Itoa(int(id))); err == nil && comment != nil {
		return comment, nil
	}

	commentResp, err := v.entRepo.GetCommentById(ctx, id)
	if err != nil {
		return nil, err
	}

	v.redisRepo.PutComment(ctx, strconv.Itoa(int(id)), commentResp)
	return commentResp, nil
}

func (v *commentUsecase) GetListCommentByVideoID(ctx context.Context, videoID uint, key string) (*models.CommentsListResponse, error) {
	//if commentCacheResp, err := v.redisRepo.GetCommentsByKey(ctx, key); err == nil && commentCacheResp.Comments != nil {
	//	return &commentCacheResp, nil
	//}

	commentResp, err := v.entRepo.GetListCommentByVideoID(ctx, videoID)
	if err != nil {
		return nil, err
	}

	//if err = v.redisRepo.PutComments(ctx, key, *commentResp); err != nil {
	//	return nil, err
	//}
	return commentResp, nil
}
