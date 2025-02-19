package repositories

import (
	"context"
	"time"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/reactions"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/repositories"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type reactionRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewReactionRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IReactionRepositoryWriter {
	return &reactionRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *reactionRepository) GetListReactionByVideoID(ctx context.Context, videoID uint) ([]*ent.Reactions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.Search")
	defer span.Finish()

	query := p.entClient.Reactions.Query()
	query = query.Where(
		reactions.DeletedAtIsNil(),
		reactions.VideoID(videoID),
	)

	// Execute the query to retrieve a page of comments.
	listReactions, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	return listReactions, nil
}

func (p *reactionRepository) CreateInBulk(ctx context.Context, reactions []*ent.Reactions) error {
	// Create a transaction
	tx, err := p.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a list of records to insert in bulk
	var reactionsBulk []*ent.ReactionsCreate
	for _, reaction := range reactions {
		reactionCreate := tx.Reactions.Create().
			SetVideoID(reaction.VideoID).
			SetDescription(reaction.Description).
			SetName(reaction.Name).
			SetNumber(reaction.Number).
			SetTimePoint(reaction.TimePoint).
			SetUpdatedAt(time.Now())
		reactionsBulk = append(reactionsBulk, reactionCreate)
	}

	// Insert the records in bulk
	if err = tx.Reactions.CreateBulk(reactionsBulk...).Exec(ctx); err != nil {
		return err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *reactionRepository) UpdateReaction(ctx context.Context, reaction *ent.Reactions) (*ent.Reactions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "reactionRepository.UpdateReaction")
	defer span.Finish()

	createdReaction, err := p.entClient.Reactions.
		UpdateOneID(reaction.ID).
		SetVideoID(reaction.VideoID).
		SetDescription(reaction.Description).
		SetName(reaction.Name).
		SetNumber(reaction.Number).
		SetTimePoint(reaction.TimePoint).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.UpdateReaction")
	}

	return createdReaction, nil
}

func (p *reactionRepository) DeleteReaction(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "reactionRepository.DeleteReaction")
	defer span.Finish()

	err := p.entClient.Reactions.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "db.DeleteReaction")
	}

	return nil
}

func (p *reactionRepository) CreateReaction(ctx context.Context, reaction *ent.Reactions) (*ent.Reactions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "reactionRepository.CreateReaction")
	defer span.Finish()

	createdReaction, err := p.entClient.Reactions.
		Create().
		SetVideoID(reaction.VideoID).
		SetDescription(reaction.Description).
		SetName(reaction.Name).
		SetNumber(reaction.Number).
		SetTimePoint(reaction.TimePoint).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.createdReaction")
	}

	return createdReaction, nil
}

func (p *reactionRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
