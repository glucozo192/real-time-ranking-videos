package repositories

import (
	"context"
	"time"

	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/videos"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/repositories"
)

type videoRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewVideoRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IVideoRepositoryWriter {
	return &videoRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *videoRepository) GetMaxVersion(ctx context.Context) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "videoRepository.GetMaxVersion")
	defer span.Finish()

	var v []struct {
		Max int64
	}
	err := p.entClient.Videos.Query().
		Aggregate(
			ent.Max(videos.FieldVersion),
		).
		Scan(ctx, &v)
	if err != nil {
		p.traceErr(span, err)
		return 0, errors.Wrap(err, "db.GetVideoById")
	}

	return v[0].Max, nil
}

func (p *videoRepository) CreateInBulk(ctx context.Context, videos []*ent.Videos) error {
	// Create a transaction
	tx, err := p.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a list of records to insert in bulk
	var videosBulk []*ent.VideosCreate
	for _, video := range videos {
		videoCreate := tx.Videos.Create().
			SetName(video.Name).
			SetDescription(video.Description).
			SetVideoURL(video.VideoURL).
			SetConfig(video.Config).
			SetPathResource(video.PathResource).
			SetLevelSystem(video.LevelSystem).
			SetStatus(video.Status).
			SetNote(video.Note).
			SetAssign(video.Assign).
			SetAuthor(video.Author).
			SetCreatedAt(time.Now())
		videosBulk = append(videosBulk, videoCreate)
	}

	// Insert the records in bulk
	if err = tx.Videos.CreateBulk(videosBulk...).Exec(ctx); err != nil {
		return err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *videoRepository) GetVideoById(ctx context.Context, id uint) (*ent.Videos, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "videoRepository.GetProductById")
	defer span.Finish()

	videoDB, err := p.entClient.Videos.Query().Where(
		videos.DeletedAtIsNil(),
		videos.ID(id),
	).Only(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.GetVideoById")
	}

	return videoDB, nil
}

func (p *videoRepository) UpdateVideo(ctx context.Context, video *ent.Videos) (*ent.Videos, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "videoRepository.UpdateVideo")
	defer span.Finish()

	videoDB, err := p.entClient.Videos.Query().Where(
		videos.DeletedAtIsNil(),
		videos.ID(video.ID),
	).Only(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.UpdateVideo")
	}

	createdVideo, err := p.entClient.Videos.
		UpdateOneID(video.ID).
		SetName(video.Name).
		SetDescription(video.Description).
		SetVideoURL(video.VideoURL).
		SetConfig(video.Config).
		SetPathResource(video.PathResource).
		SetLevelSystem(video.LevelSystem).
		SetStatus(video.Status).
		SetNote(video.Note).
		SetAssign(video.Assign).
		SetAuthor(video.Author).
		SetUpdatedAt(time.Now()).
		SetVersion(videoDB.Version + 1).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.UpdateVideo")
	}

	return createdVideo, nil
}

func (p *videoRepository) DeleteVideo(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "videoRepository.DeleteVideo")
	defer span.Finish()

	err := p.entClient.Videos.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "db.DeleteVideo")
	}

	return nil
}

func (p *videoRepository) CreateVideo(ctx context.Context, video *ent.Videos) (*ent.Videos, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "videoRepository.CreateVideo")
	defer span.Finish()

	createdVideo, err := p.entClient.Videos.
		Create().
		SetName(video.Name).
		SetDescription(video.Description).
		SetVideoURL(video.VideoURL).
		SetConfig(video.Config).
		SetPathResource(video.PathResource).
		SetLevelSystem(video.LevelSystem).
		SetStatus(video.Status).
		SetNote(video.Note).
		SetAssign(video.Assign).
		SetAuthor(video.Author).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.createdVideo")
	}

	return createdVideo, nil
}

func (p *videoRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
