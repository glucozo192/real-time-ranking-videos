package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/viewers"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type viewerRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewViewerRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IViewerRepositoryReader {
	return &viewerRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *viewerRepository) GetViewerById(ctx context.Context, id uint) (*ent.Viewers, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "viewerRepository.GetProductById")
	defer span.Finish()

	viewerDB, err := p.entClient.Viewers.Query().Where(
		viewers.DeletedAtIsNil(),
		viewers.ID(id),
	).Only(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.GetViewerById")
	}

	return viewerDB, nil
}

func (p *viewerRepository) GetListViewerByVideoID(ctx context.Context, videoID uint) (*models.ViewerListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "viewerRepository.Search")
	defer span.Finish()

	query := p.entClient.Viewers.Query()
	query = query.Where(
		viewers.DeletedAtIsNil(),
		viewers.VideoID(videoID),
	)

	// Execute the query to count total records.
	totalCount, err := query.Count(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.Count")
	}

	// Execute the query to retrieve a page of viewers.
	listViewers, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	// Create the response.
	response := &models.ViewerListResponse{
		TotalCount: totalCount,
		TotalPages: 1,
		Page:       1,
		Size:       totalCount,
		HasMore:    false,
		Viewers:    listViewers,
	}

	return response, nil

}

func (p *viewerRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
