package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/comments"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type commentRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewCommentRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.ICommentRepositoryReader {
	return &commentRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *commentRepository) GetCommentById(ctx context.Context, id uint) (*ent.Comments, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.GetProductById")
	defer span.Finish()

	commentDB, err := p.entClient.Comments.Query().Where(
		comments.DeletedAtIsNil(),
		comments.ID(id),
	).Only(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.GetCommentById")
	}

	return commentDB, nil
}

func (p *commentRepository) GetListCommentByVideoID(ctx context.Context, videoID uint) (*models.CommentsListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.Search")
	defer span.Finish()

	query := p.entClient.Comments.Query()
	query = query.Where(
		comments.DeletedAtIsNil(),
		comments.VideoID(videoID),
	)

	// Execute the query to count total records.
	totalCount, err := query.Count(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.Count")
	}

	// Execute the query to retrieve a page of comments.
	listComments, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	// Create the response.
	response := &models.CommentsListResponse{
		TotalCount: totalCount,
		TotalPages: 1,
		Page:       1,
		Size:       totalCount,
		HasMore:    false,
		Comments:   listComments,
	}

	return response, nil

}

func (p *commentRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
