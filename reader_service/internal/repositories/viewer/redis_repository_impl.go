package repositories

import (
	"context"
	"encoding/json"

	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/models"

	"github.com/glu/video-real-time-ranking/core/pkg/logger"
	"github.com/glu/video-real-time-ranking/ent"
	"github.com/glu/video-real-time-ranking/reader_service/config"
	"github.com/glu/video-real-time-ranking/reader_service/internal/domain/repositories"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

const (
	redisViewerPrefixKey = "reader:viewer"
)

type redisViewerRepository struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisViewerRepository(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) repositories.IViewerCacheRepository {
	return &redisViewerRepository{log: log, cfg: cfg, redisClient: redisClient}
}

func (r *redisViewerRepository) PutViewer(ctx context.Context, key string, viewer *ent.Viewers) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisViewerRepository.PutViewer")
	defer span.Finish()

	viewerBytes, err := json.Marshal(viewer)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisViewerPrefixKey(), key, viewerBytes).Err(); err != nil {
		r.log.WarnMsg("redisClient.HSetNX", err)
		return
	}
	r.log.Debugf("HSetNX prefix: %s, key: %s", r.getRedisViewerPrefixKey(), key)
}

func (r *redisViewerRepository) PutViewers(ctx context.Context, key string, viewers models.ViewerListResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisViewerRepository.PutViewers")
	defer span.Finish()

	// Serialize the ViewersListResponse to JSON.
	viewerJSON, err := json.Marshal(viewers)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return errors.Wrap(err, "json.Marshal")
	}

	// Store the JSON string in Redis under the specified key.
	err = r.redisClient.Set(ctx, key, viewerJSON, 0).Err()
	if err != nil {
		r.log.WarnMsg("redisClient.Set", err)
		return errors.Wrap(err, "redisClient.Set")
	}

	r.log.Debugf("Set key: %s in Redis", key)
	return nil
}

func (r *redisViewerRepository) GetViewer(ctx context.Context, key string) (*ent.Viewers, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisViewerRepository.GetViewer")
	defer span.Finish()

	viewerBytes, err := r.redisClient.HGet(ctx, r.getRedisViewerPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	viewer := &ent.Viewers{}
	if err := json.Unmarshal(viewerBytes, &viewer); err != nil {
		return nil, err
	}

	r.log.Debugf("HGet prefix: %s, key: %s", r.getRedisViewerPrefixKey(), key)
	return viewer, nil
}

func (r *redisViewerRepository) GetViewersByKey(ctx context.Context, key string) (models.ViewerListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisViewerRepository.GetViewers")
	defer span.Finish()

	// Retrieve the JSON string from Redis by key.
	viewerJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return models.ViewerListResponse{}, nil // Key not found, return an empty result.
		}
		r.log.WarnMsg("redisClient.Get", err)
		return models.ViewerListResponse{}, errors.Wrap(err, "redisClient.Get")
	}

	// Deserialize the JSON string into a ViewerListResponse.
	var viewers models.ViewerListResponse
	if err := json.Unmarshal([]byte(viewerJSON), &viewers); err != nil {
		r.log.WarnMsg("json.Unmarshal", err)
		return models.ViewerListResponse{}, errors.Wrap(err, "json.Unmarshal")
	}

	return viewers, nil
}

func (r *redisViewerRepository) GetViewers(ctx context.Context) ([]*ent.Viewers, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisViewerRepository.GetViewers")
	defer span.Finish()

	// Use HGetAll to get all viewer entries.
	viewerMap, err := r.redisClient.HGetAll(ctx, r.getRedisViewerPrefixKey()).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGetAll", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGetAll")
	}

	// Iterate over the map and unmarshal each entry into a viewer.
	viewers := make([]*ent.Viewers, 0, len(viewerMap))
	for _, viewerJSON := range viewerMap {
		viewer := &ent.Viewers{}
		if err := json.Unmarshal([]byte(viewerJSON), &viewer); err != nil {
			return nil, err
		}
		viewers = append(viewers, viewer)
	}

	r.log.Debugf("HGetAll prefix: %s", r.getRedisViewerPrefixKey())
	return viewers, nil
}

func (r *redisViewerRepository) DelViewer(ctx context.Context, key string) {
	if err := r.redisClient.HDel(ctx, r.getRedisViewerPrefixKey(), key).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("HDel prefix: %s, key: %s", r.getRedisViewerPrefixKey(), key)
}

func (r *redisViewerRepository) DelAllViewers(ctx context.Context) {
	if err := r.redisClient.Del(ctx, r.getRedisViewerPrefixKey()).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("Del key: %s", r.getRedisViewerPrefixKey())
}

func (r *redisViewerRepository) getRedisViewerPrefixKey() string {
	if r.cfg.ServiceSettings.RedisViewerPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisViewerPrefixKey
	}

	return redisViewerPrefixKey
}
