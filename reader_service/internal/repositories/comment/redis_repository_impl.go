package repositories

import (
	"context"
	"encoding/json"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

const (
	redisCommentPrefixKey = "reader:comment"
)

type redisCommentRepository struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisCommentRepository(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) repositories.ICommentCacheRepository {
	return &redisCommentRepository{log: log, cfg: cfg, redisClient: redisClient}
}

func (r *redisCommentRepository) PutComment(ctx context.Context, key string, comment *ent.Comments) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisCommentRepository.PutComment")
	defer span.Finish()

	commentBytes, err := json.Marshal(comment)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisCommentPrefixKey(), key, commentBytes).Err(); err != nil {
		r.log.WarnMsg("redisClient.HSetNX", err)
		return
	}
	r.log.Debugf("HSetNX prefix: %s, key: %s", r.getRedisCommentPrefixKey(), key)
}

func (r *redisCommentRepository) PutComments(ctx context.Context, key string, comments models.CommentsListResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisCommentRepository.PutComments")
	defer span.Finish()

	// Serialize the CommentsListResponse to JSON.
	commentJSON, err := json.Marshal(comments)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return errors.Wrap(err, "json.Marshal")
	}

	// Store the JSON string in Redis under the specified key.
	err = r.redisClient.Set(ctx, key, commentJSON, 0).Err()
	if err != nil {
		r.log.WarnMsg("redisClient.Set", err)
		return errors.Wrap(err, "redisClient.Set")
	}

	r.log.Debugf("Set key: %s in Redis", key)
	return nil
}

func (r *redisCommentRepository) GetComment(ctx context.Context, key string) (*ent.Comments, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisCommentRepository.GetComment")
	defer span.Finish()

	commentBytes, err := r.redisClient.HGet(ctx, r.getRedisCommentPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	comment := &ent.Comments{}
	if err := json.Unmarshal(commentBytes, &comment); err != nil {
		return nil, err
	}

	r.log.Debugf("HGet prefix: %s, key: %s", r.getRedisCommentPrefixKey(), key)
	return comment, nil
}

func (r *redisCommentRepository) GetCommentsByKey(ctx context.Context, key string) (models.CommentsListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisCommentRepository.GetComments")
	defer span.Finish()

	// Retrieve the JSON string from Redis by key.
	commentJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return models.CommentsListResponse{}, nil // Key not found, return an empty result.
		}
		r.log.WarnMsg("redisClient.Get", err)
		return models.CommentsListResponse{}, errors.Wrap(err, "redisClient.Get")
	}

	// Deserialize the JSON string into a CommentsListResponse.
	var comments models.CommentsListResponse
	if err := json.Unmarshal([]byte(commentJSON), &comments); err != nil {
		r.log.WarnMsg("json.Unmarshal", err)
		return models.CommentsListResponse{}, errors.Wrap(err, "json.Unmarshal")
	}

	return comments, nil
}

func (r *redisCommentRepository) GetComments(ctx context.Context) ([]*ent.Comments, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisCommentRepository.GetComments")
	defer span.Finish()

	// Use HGetAll to get all comment entries.
	commentMap, err := r.redisClient.HGetAll(ctx, r.getRedisCommentPrefixKey()).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGetAll", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGetAll")
	}

	// Iterate over the map and unmarshal each entry into a comment.
	comments := make([]*ent.Comments, 0, len(commentMap))
	for _, commentJSON := range commentMap {
		comment := &ent.Comments{}
		if err := json.Unmarshal([]byte(commentJSON), &comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	r.log.Debugf("HGetAll prefix: %s", r.getRedisCommentPrefixKey())
	return comments, nil
}

func (r *redisCommentRepository) DelComment(ctx context.Context, key string) {
	if err := r.redisClient.HDel(ctx, r.getRedisCommentPrefixKey(), key).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("HDel prefix: %s, key: %s", r.getRedisCommentPrefixKey(), key)
}

func (r *redisCommentRepository) DelAllComments(ctx context.Context) {
	if err := r.redisClient.Del(ctx, r.getRedisCommentPrefixKey()).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("Del key: %s", r.getRedisCommentPrefixKey())
}

func (r *redisCommentRepository) getRedisCommentPrefixKey() string {
	if r.cfg.ServiceSettings.RedisCommentPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisCommentPrefixKey
	}

	return redisCommentPrefixKey
}
