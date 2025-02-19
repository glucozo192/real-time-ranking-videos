package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/reactions"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type reactionRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewReactionRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IReactionRepositoryReader {
	return &reactionRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *reactionRepository) GetReactionById(ctx context.Context, id uint) (*ent.Reactions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "reactionRepository.GetProductById")
	defer span.Finish()

	reactionDB, err := p.entClient.Reactions.Query().Where(
		reactions.DeletedAtIsNil(),
		reactions.ID(id),
	).Only(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.GetReactionById")
	}

	return reactionDB, nil
}

func (p *reactionRepository) GetListReactionByVideoID(ctx context.Context, videoID uint) (*models.ReactionsListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "reactionRepository.Search")
	defer span.Finish()

	query := p.entClient.Reactions.Query()
	query = query.Where(
		reactions.DeletedAtIsNil(),
		reactions.VideoID(videoID),
	)

	// Execute the query to count total records.
	totalCount, err := query.Count(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.Count")
	}

	// Execute the query to retrieve a page of reactions.
	listReactions, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	// Create the response.
	response := &models.ReactionsListResponse{
		TotalCount: totalCount,
		TotalPages: 1,
		Page:       1,
		Size:       totalCount,
		HasMore:    false,
		Reactions:  listReactions,
	}

	return response, nil

}

func (p *reactionRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
