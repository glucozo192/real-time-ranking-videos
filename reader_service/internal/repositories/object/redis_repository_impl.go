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
	redisObjectPrefixKey = "reader:object"
)

type redisObjectRepository struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisObjectRepository(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) repositories.IObjectCacheRepository {
	return &redisObjectRepository{log: log, cfg: cfg, redisClient: redisClient}
}

func (r *redisObjectRepository) PutObject(ctx context.Context, key string, object *ent.Objects) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisObjectRepository.PutObject")
	defer span.Finish()

	objectBytes, err := json.Marshal(object)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisObjectPrefixKey(), key, objectBytes).Err(); err != nil {
		r.log.WarnMsg("redisClient.HSetNX", err)
		return
	}
	r.log.Debugf("HSetNX prefix: %s, key: %s", r.getRedisObjectPrefixKey(), key)
}

func (r *redisObjectRepository) PutObjects(ctx context.Context, key string, objects models.ObjectsListResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisObjectRepository.PutObjects")
	defer span.Finish()

	// Serialize the ObjectsListResponse to JSON.
	objectJSON, err := json.Marshal(objects)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return errors.Wrap(err, "json.Marshal")
	}

	// Store the JSON string in Redis under the specified key.
	err = r.redisClient.Set(ctx, key, objectJSON, 0).Err()
	if err != nil {
		r.log.WarnMsg("redisClient.Set", err)
		return errors.Wrap(err, "redisClient.Set")
	}

	r.log.Debugf("Set key: %s in Redis", key)
	return nil
}

func (r *redisObjectRepository) GetObject(ctx context.Context, key string) (*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisObjectRepository.GetObject")
	defer span.Finish()

	objectBytes, err := r.redisClient.HGet(ctx, r.getRedisObjectPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	object := &ent.Objects{}
	if err := json.Unmarshal(objectBytes, &object); err != nil {
		return nil, err
	}

	r.log.Debugf("HGet prefix: %s, key: %s", r.getRedisObjectPrefixKey(), key)
	return object, nil
}

func (r *redisObjectRepository) GetObjectsByKey(ctx context.Context, key string) (models.ObjectsListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisObjectRepository.GetObjects")
	defer span.Finish()

	// Retrieve the JSON string from Redis by key.
	objectJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return models.ObjectsListResponse{}, nil // Key not found, return an empty result.
		}
		r.log.WarnMsg("redisClient.Get", err)
		return models.ObjectsListResponse{}, errors.Wrap(err, "redisClient.Get")
	}

	// Deserialize the JSON string into a ObjectsListResponse.
	var objects models.ObjectsListResponse
	if err := json.Unmarshal([]byte(objectJSON), &objects); err != nil {
		r.log.WarnMsg("json.Unmarshal", err)
		return models.ObjectsListResponse{}, errors.Wrap(err, "json.Unmarshal")
	}

	return objects, nil
}

func (r *redisObjectRepository) GetObjects(ctx context.Context) ([]*ent.Objects, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisObjectRepository.GetObjects")
	defer span.Finish()

	// Use HGetAll to get all object entries.
	objectMap, err := r.redisClient.HGetAll(ctx, r.getRedisObjectPrefixKey()).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGetAll", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGetAll")
	}

	// Iterate over the map and unmarshal each entry into a object.
	objects := make([]*ent.Objects, 0, len(objectMap))
	for _, objectJSON := range objectMap {
		object := &ent.Objects{}
		if err := json.Unmarshal([]byte(objectJSON), &object); err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}

	r.log.Debugf("HGetAll prefix: %s", r.getRedisObjectPrefixKey())
	return objects, nil
}

func (r *redisObjectRepository) DelObject(ctx context.Context, key string) {
	if err := r.redisClient.HDel(ctx, r.getRedisObjectPrefixKey(), key).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("HDel prefix: %s, key: %s", r.getRedisObjectPrefixKey(), key)
}

func (r *redisObjectRepository) DelAllObjects(ctx context.Context) {
	if err := r.redisClient.Del(ctx, r.getRedisObjectPrefixKey()).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("Del key: %s", r.getRedisObjectPrefixKey())
}

func (r *redisObjectRepository) getRedisObjectPrefixKey() string {
	if r.cfg.ServiceSettings.RedisObjectPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisObjectPrefixKey
	}

	return redisObjectPrefixKey
}
