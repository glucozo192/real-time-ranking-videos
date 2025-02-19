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
	redisVideoPrefixKey = "reader:video"
)

type redisVideoRepository struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisVideoRepository(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) repositories.IVideoCacheRepository {
	return &redisVideoRepository{log: log, cfg: cfg, redisClient: redisClient}
}

func (r *redisVideoRepository) PutVideo(ctx context.Context, key string, video *ent.Videos) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisVideoRepository.PutVideo")
	defer span.Finish()

	videoBytes, err := json.Marshal(video)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisVideoPrefixKey(), key, videoBytes).Err(); err != nil {
		r.log.WarnMsg("redisClient.HSetNX", err)
		return
	}
	r.log.Debugf("HSetNX prefix: %s, key: %s", r.getRedisVideoPrefixKey(), key)
}

func (r *redisVideoRepository) PutVideos(ctx context.Context, key string, videos models.VideosListResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisVideoRepository.PutVideos")
	defer span.Finish()

	// Serialize the VideosListResponse to JSON.
	videoJSON, err := json.Marshal(videos)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return errors.Wrap(err, "json.Marshal")
	}

	// Store the JSON string in Redis under the specified key.
	err = r.redisClient.Set(ctx, key, videoJSON, 0).Err()
	if err != nil {
		r.log.WarnMsg("redisClient.Set", err)
		return errors.Wrap(err, "redisClient.Set")
	}

	r.log.Debugf("Set key: %s in Redis", key)
	return nil
}

func (r *redisVideoRepository) GetVideo(ctx context.Context, key string) (*ent.Videos, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisVideoRepository.GetVideo")
	defer span.Finish()

	videoBytes, err := r.redisClient.HGet(ctx, r.getRedisVideoPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	video := &ent.Videos{}
	if err := json.Unmarshal(videoBytes, &video); err != nil {
		return nil, err
	}

	r.log.Debugf("HGet prefix: %s, key: %s", r.getRedisVideoPrefixKey(), key)
	return video, nil
}

func (r *redisVideoRepository) GetVideosByKey(ctx context.Context, key string) (models.VideosListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisVideoRepository.GetVideos")
	defer span.Finish()

	// Retrieve the JSON string from Redis by key.
	videoJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return models.VideosListResponse{}, nil // Key not found, return an empty result.
		}
		r.log.WarnMsg("redisClient.Get", err)
		return models.VideosListResponse{}, errors.Wrap(err, "redisClient.Get")
	}

	// Deserialize the JSON string into a VideosListResponse.
	var videos models.VideosListResponse
	if err := json.Unmarshal([]byte(videoJSON), &videos); err != nil {
		r.log.WarnMsg("json.Unmarshal", err)
		return models.VideosListResponse{}, errors.Wrap(err, "json.Unmarshal")
	}

	return videos, nil
}

func (r *redisVideoRepository) GetVideos(ctx context.Context) ([]*ent.Videos, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisVideoRepository.GetVideos")
	defer span.Finish()

	// Use HGetAll to get all video entries.
	videoMap, err := r.redisClient.HGetAll(ctx, r.getRedisVideoPrefixKey()).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGetAll", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGetAll")
	}

	// Iterate over the map and unmarshal each entry into a video.
	videos := make([]*ent.Videos, 0, len(videoMap))
	for _, videoJSON := range videoMap {
		video := &ent.Videos{}
		if err := json.Unmarshal([]byte(videoJSON), &video); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	r.log.Debugf("HGetAll prefix: %s", r.getRedisVideoPrefixKey())
	return videos, nil
}

func (r *redisVideoRepository) DelVideo(ctx context.Context, key string) {
	if err := r.redisClient.HDel(ctx, r.getRedisVideoPrefixKey(), key).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("HDel prefix: %s, key: %s", r.getRedisVideoPrefixKey(), key)
}

func (r *redisVideoRepository) DelAllVideos(ctx context.Context) {
	if err := r.redisClient.Del(ctx, r.getRedisVideoPrefixKey()).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("Del key: %s", r.getRedisVideoPrefixKey())
}

func (r *redisVideoRepository) getRedisVideoPrefixKey() string {
	if r.cfg.ServiceSettings.RedisVideoPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisVideoPrefixKey
	}

	return redisVideoPrefixKey
}
