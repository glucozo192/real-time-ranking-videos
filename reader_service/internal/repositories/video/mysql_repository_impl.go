package repositories

import (
	"context"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/ent/videos"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type videoRepository struct {
	log       logger.Logger
	cfg       *config.Config
	entClient *ent.Client
}

func NewVideoRepository(log logger.Logger, cfg *config.Config, entClient *ent.Client) repositories.IVideoRepositoryReader {
	return &videoRepository{log: log, cfg: cfg, entClient: entClient}
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

func (p *videoRepository) GetVideoByVideoUrl(ctx context.Context, url string) (*ent.Videos, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "videoRepository.VideoByUrl")
	defer span.Finish()

	videoDB, err := p.entClient.Videos.Query().Where(
		videos.DeletedAtIsNil(),
		videos.VideoURL(url),
	).Only(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.GetVideoByUrl")
	}

	return videoDB, nil
}

func (p *videoRepository) SearchVideoByParams(ctx context.Context, req models.SearchVideoRequest) (*models.VideosListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "videoRepository.Search")
	defer span.Finish()

	query := p.entClient.Videos.Query()
	query = query.Where(videos.DeletedAtIsNil())

	// Add filters based on the request parameters.
	if req.ID != nil {
		query = query.Where(videos.ID(uint(*req.ID)))
	}
	if req.Name != nil && *req.Name != "" {
		query = query.Where(videos.NameContains(*req.Name))
	}
	if req.Author != nil && *req.Author != "" {
		query = query.Where(videos.Author(*req.Author))
	}
	if req.Assign != nil && *req.Assign != "" {
		query = query.Where(videos.Assign(*req.Assign))
	}
	if req.LevelSystem != nil && *req.LevelSystem != "" {
		query = query.Where(videos.LevelSystem(*req.LevelSystem))
	}

	// Execute the query to count total records.
	totalCount, err := query.Count(ctx)
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "db.Count")
	}

	// Apply pagination.
	if req.Pagination.Size > 0 {
		query = query.Limit(req.Pagination.Size)
	}
	if req.Pagination.Page > 0 {
		offset := (req.Pagination.Page - 1) * req.Pagination.Size
		query = query.Offset(offset)
	}
	//if req.Pagination.OrderBy != "" {
	query = query.Order(ent.Desc("id"))
	//}

	// Execute the query to retrieve a page of videos.
	listVideos, err := query.All(context.Background())
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "query.All")
	}

	// Calculate pagination details.
	totalPages := (totalCount + req.Pagination.Size - 1) / req.Pagination.Size
	hasMore := (req.Pagination.Page < totalPages)

	// Create the response.
	response := &models.VideosListResponse{
		TotalCount: totalCount,
		TotalPages: totalPages,
		Page:       req.Pagination.Page,
		Size:       req.Pagination.Size,
		HasMore:    hasMore,
		Videos:     listVideos,
	}

	return response, nil

}

func (p *videoRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
