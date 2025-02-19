package repositories

import (
	"context"
	"time"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/viewers"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/repositories"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type viewerRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewViewerRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IViewerRepositoryWriter {
	return &viewerRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *viewerRepository) GetListViewerByVideoID(ctx context.Context, videoID uint) ([]*ent.Viewers, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.Search")
	defer span.Finish()

	query := p.entClient.Viewers.Query()
	query = query.Where(
		viewers.DeletedAtIsNil(),
		viewers.VideoID(videoID),
	)

	// Execute the query to retrieve a page of comments.
	listViewers, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	return listViewers, nil
}

func (p *viewerRepository) CreateInBulk(ctx context.Context, viewers []*ent.Viewers) error {
	// Create a transaction
	tx, err := p.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a list of records to insert in bulk
	var viewersBulk []*ent.ViewersCreate
	for _, viewer := range viewers {
		viewerCreate := tx.Viewers.Create().
			SetVideoID(viewer.VideoID).
			SetNumber(viewer.Number).
			SetTimePoint(viewer.TimePoint).
			SetUpdatedAt(time.Now())
		viewersBulk = append(viewersBulk, viewerCreate)
	}

	// Insert the records in bulk
	if err = tx.Viewers.CreateBulk(viewersBulk...).Exec(ctx); err != nil {
		return err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *viewerRepository) UpdateViewer(ctx context.Context, viewer *ent.Viewers) (*ent.Viewers, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "viewerRepository.UpdateViewer")
	defer span.Finish()

	createdViewer, err := p.entClient.Viewers.
		UpdateOneID(viewer.ID).
		SetVideoID(viewer.VideoID).
		SetNumber(viewer.Number).
		SetTimePoint(viewer.TimePoint).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.UpdateViewer")
	}

	return createdViewer, nil
}

func (p *viewerRepository) DeleteViewer(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "viewerRepository.DeleteViewer")
	defer span.Finish()

	err := p.entClient.Viewers.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "db.DeleteViewer")
	}

	return nil
}

func (p *viewerRepository) CreateViewer(ctx context.Context, viewer *ent.Viewers) (*ent.Viewers, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "viewerRepository.CreateViewer")
	defer span.Finish()

	createdViewer, err := p.entClient.Viewers.
		Create().
		SetVideoID(viewer.VideoID).
		SetNumber(viewer.Number).
		SetTimePoint(viewer.TimePoint).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.createdViewer")
	}

	return createdViewer, nil
}

func (p *viewerRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
