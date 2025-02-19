package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/objects"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type objectRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewObjectRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IObjectRepositoryReader {
	return &objectRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *objectRepository) GetObjectById(ctx context.Context, id uint) (*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "objectRepository.GetProductById")
	defer span.Finish()

	objectDB, err := p.entClient.Objects.Query().Where(
		objects.DeletedAtIsNil(),
		objects.ID(id),
	).Only(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.GetObjectById")
	}

	return objectDB, nil
}

func (p *objectRepository) GetListObjectByVideoID(ctx context.Context, videoID uint) (*models.ObjectsListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "objectRepository.Search")
	defer span.Finish()

	query := p.entClient.Objects.Query()
	query = query.Where(
		objects.DeletedAtIsNil(),
		objects.VideoID(videoID),
	)

	// Execute the query to count total records.
	totalCount, err := query.Count(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.Count")
	}

	// Execute the query to retrieve a page of objects.
	listObjects, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	// Create the response.
	response := &models.ObjectsListResponse{
		TotalCount: totalCount,
		TotalPages: 1,
		Page:       1,
		Size:       totalCount,
		HasMore:    false,
		Objects:    listObjects,
	}

	return response, nil

}

func (p *objectRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}

func (p *objectRepository) GetListObjectByVideoIDV2(ctx context.Context, videoID uint) ([]*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "objectRepository.Search")
	defer span.Finish()

	query := p.entClient.Objects.Query()
	query = query.Where(
		objects.DeletedAtIsNil(),
		objects.VideoID(videoID),
	)

	listObjects, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	return listObjects, nil

}
