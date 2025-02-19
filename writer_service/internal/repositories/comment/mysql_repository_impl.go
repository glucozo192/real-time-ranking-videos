package repositories

import (
	"context"
	"time"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/comments"
	"github.com/glu/video-real-time-ranking/writer_service/config"
	"github.com/glu/video-real-time-ranking/writer_service/internal/domain/repositories"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type commentRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewCommentRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.ICommentRepositoryWriter {
	return &commentRepository{log: log, cfg: cfg, entClient: entClient}
}

func (p *commentRepository) GetListCommentByVideoID(ctx context.Context, videoID uint) ([]*ent.Comments, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.Search")
	defer span.Finish()

	query := p.entClient.Comments.Query()
	query = query.Where(
		comments.DeletedAtIsNil(),
		comments.VideoID(videoID),
	)

	// Execute the query to retrieve a page of comments.
	listComments, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	return listComments, nil
}

func (p *commentRepository) CreateInBulk(ctx context.Context, comments []*ent.Comments) error {
	// Create a transaction
	tx, err := p.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a list of records to insert in bulk
	var commentsBulk []*ent.CommentsCreate
	for _, comment := range comments {
		commentCreate := tx.Comments.Create().
			SetVideoID(comment.VideoID).
			SetDescription(comment.Description).
			SetComment(comment.Comment).
			SetUserName(comment.UserName).
			SetAvatar(comment.Avatar).
			SetTimePoint(comment.TimePoint).
			SetUpdatedAt(time.Now())
		commentsBulk = append(commentsBulk, commentCreate)
	}

	// Insert the records in bulk
	if err = tx.Comments.CreateBulk(commentsBulk...).Exec(ctx); err != nil {
		return err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *commentRepository) UpdateComment(ctx context.Context, comment *ent.Comments) (*ent.Comments, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.UpdateComment")
	defer span.Finish()

	createdComment, err := p.entClient.Comments.
		UpdateOneID(comment.ID).
		SetVideoID(comment.VideoID).
		SetDescription(comment.Description).
		SetComment(comment.Comment).
		SetUserName(comment.UserName).
		SetAvatar(comment.Avatar).
		SetTimePoint(comment.TimePoint).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.UpdateComment")
	}

	return createdComment, nil
}

func (p *commentRepository) DeleteComment(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.DeleteComment")
	defer span.Finish()

	err := p.entClient.Comments.
		UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "db.DeleteComment")
	}

	return nil
}

func (p *commentRepository) CreateComment(ctx context.Context, comment *ent.Comments) (*ent.Comments, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "commentRepository.CreateComment")
	defer span.Finish()

	createdComment, err := p.entClient.Comments.
		Create().
		SetVideoID(comment.VideoID).
		SetDescription(comment.Description).
		SetComment(comment.Comment).
		SetUserName(comment.UserName).
		SetAvatar(comment.Avatar).
		SetTimePoint(comment.TimePoint).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.createdComment")
	}

	return createdComment, nil
}

func (p *commentRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
