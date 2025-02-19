package repositories

import (
	"context"
	"time"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/objects"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/repositories"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type objectRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewObjectRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IObjectRepositoryWriter {
	return &objectRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *objectRepository) GetListObjectByVideoID(ctx context.Context, videoID uint) ([]*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.Search")
	defer span.Finish()

	query := p.entClient.Objects.Query()
	query = query.Where(
		objects.DeletedAtIsNil(),
		objects.VideoID(videoID),
	)

	// Execute the query to retrieve a page of comments.
	listObjects, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	return listObjects, nil
}

func (p *objectRepository) UpdateObject(ctx context.Context, object *ent.Objects) (*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "objectRepository.UpdateObject")
	defer span.Finish()

	createdObject, err := p.entClient.Objects.
		UpdateOneID(object.ID).
		SetVideoID(object.VideoID).
		SetDescription(object.Description).
		SetCoordinateX(object.CoordinateX).
		SetCoordinateY(object.CoordinateY).
		SetLength(object.Length).
		SetWidth(object.Width).
		SetOrder(object.Order).
		SetTimeStart(object.TimeStart).
		SetTimeEnd(object.TimeEnd).
		SetUpdatedAt(time.Now()).
		SetTouchVector(object.TouchVector).
		SetMarkerName(object.MarkerName).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.UpdateObject")
	}

	return createdObject, nil
}

func (p *objectRepository) DeleteObject(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "objectRepository.DeleteObject")
	defer span.Finish()

	err := p.entClient.Objects.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "db.DeleteObject")
	}

	return nil
}

func (p *objectRepository) CreateObject(ctx context.Context, object *ent.Objects) (*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "objectRepository.CreateObject")
	defer span.Finish()

	createdObject, err := p.entClient.Objects.
		Create().
		SetVideoID(object.VideoID).
		SetDescription(object.Description).
		SetCoordinateX(object.CoordinateX).
		SetCoordinateY(object.CoordinateY).
		SetLength(object.Length).
		SetWidth(object.Width).
		SetOrder(object.Order).
		SetTimeStart(object.TimeStart).
		SetTimeEnd(object.TimeEnd).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.createdObject")
	}

	return createdObject, nil
}

func (p *objectRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}

func (p *objectRepository) CreateInBulk(ctx context.Context, objects []*ent.Objects) error {
	// Create a transaction
	tx, err := p.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback() // TODO: check in future

	// Create a list of records to insert in bulk
	var objectsBulk []*ent.ObjectsCreate
	for _, object := range objects {
		videoCreate := tx.Objects.Create().
			SetVideoID(object.VideoID).
			SetDescription(object.Description).
			SetCoordinateX(object.CoordinateX).
			SetCoordinateY(object.CoordinateY).
			SetLength(object.Length).
			SetWidth(object.Width).
			SetOrder(object.Order).
			SetTimeStart(object.TimeStart).
			SetTimeEnd(object.TimeEnd).
			SetUpdatedAt(time.Now()).
			SetMarkerName(object.MarkerName).
			SetTouchVector(object.TouchVector)
		objectsBulk = append(objectsBulk, videoCreate)
	}

	// Insert the records in bulk
	if err = tx.Objects.CreateBulk(objectsBulk...).Exec(ctx); err != nil {
		return err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
