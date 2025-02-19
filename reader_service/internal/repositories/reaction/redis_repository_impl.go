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
	redisReactionPrefixKey = "reader:reaction"
)

type redisReactionRepository struct {
	log         logger.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisReactionRepository(log logger.Logger, cfg *config.Config, redisClient redis.UniversalClient) repositories.IReactionCacheRepository {
	return &redisReactionRepository{log: log, cfg: cfg, redisClient: redisClient}
}

func (r *redisReactionRepository) PutReaction(ctx context.Context, key string, reaction *ent.Reactions) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisReactionRepository.PutReaction")
	defer span.Finish()

	reactionBytes, err := json.Marshal(reaction)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return
	}

	if err := r.redisClient.HSetNX(ctx, r.getRedisReactionPrefixKey(), key, reactionBytes).Err(); err != nil {
		r.log.WarnMsg("redisClient.HSetNX", err)
		return
	}
	r.log.Debugf("HSetNX prefix: %s, key: %s", r.getRedisReactionPrefixKey(), key)
}

func (r *redisReactionRepository) PutReactions(ctx context.Context, key string, reactions models.ReactionsListResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisReactionRepository.PutReactions")
	defer span.Finish()

	// Serialize the ReactionsListResponse to JSON.
	reactionJSON, err := json.Marshal(reactions)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return errors.Wrap(err, "json.Marshal")
	}

	// Store the JSON string in Redis under the specified key.
	err = r.redisClient.Set(ctx, key, reactionJSON, 0).Err()
	if err != nil {
		r.log.WarnMsg("redisClient.Set", err)
		return errors.Wrap(err, "redisClient.Set")
	}

	r.log.Debugf("Set key: %s in Redis", key)
	return nil
}

func (r *redisReactionRepository) GetReaction(ctx context.Context, key string) (*ent.Reactions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisReactionRepository.GetReaction")
	defer span.Finish()

	reactionBytes, err := r.redisClient.HGet(ctx, r.getRedisReactionPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}

	reaction := &ent.Reactions{}
	if err := json.Unmarshal(reactionBytes, &reaction); err != nil {
		return nil, err
	}

	r.log.Debugf("HGet prefix: %s, key: %s", r.getRedisReactionPrefixKey(), key)
	return reaction, nil
}

func (r *redisReactionRepository) GetReactionsByKey(ctx context.Context, key string) (models.ReactionsListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisReactionRepository.GetReactions")
	defer span.Finish()

	// Retrieve the JSON string from Redis by key.
	reactionJSON, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return models.ReactionsListResponse{}, nil // Key not found, return an empty result.
		}
		r.log.WarnMsg("redisClient.Get", err)
		return models.ReactionsListResponse{}, errors.Wrap(err, "redisClient.Get")
	}

	// Deserialize the JSON string into a ReactionsListResponse.
	var reactions models.ReactionsListResponse
	if err := json.Unmarshal([]byte(reactionJSON), &reactions); err != nil {
		r.log.WarnMsg("json.Unmarshal", err)
		return models.ReactionsListResponse{}, errors.Wrap(err, "json.Unmarshal")
	}

	return reactions, nil
}

func (r *redisReactionRepository) GetReactions(ctx context.Context) ([]*ent.Reactions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisReactionRepository.GetReactions")
	defer span.Finish()

	// Use HGetAll to get all reaction entries.
	reactionMap, err := r.redisClient.HGetAll(ctx, r.getRedisReactionPrefixKey()).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGetAll", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGetAll")
	}

	// Iterate over the map and unmarshal each entry into a reaction.
	reactions := make([]*ent.Reactions, 0, len(reactionMap))
	for _, reactionJSON := range reactionMap {
		reaction := &ent.Reactions{}
		if err := json.Unmarshal([]byte(reactionJSON), &reaction); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}

	r.log.Debugf("HGetAll prefix: %s", r.getRedisReactionPrefixKey())
	return reactions, nil
}

func (r *redisReactionRepository) DelReaction(ctx context.Context, key string) {
	if err := r.redisClient.HDel(ctx, r.getRedisReactionPrefixKey(), key).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("HDel prefix: %s, key: %s", r.getRedisReactionPrefixKey(), key)
}

func (r *redisReactionRepository) DelAllReactions(ctx context.Context) {
	if err := r.redisClient.Del(ctx, r.getRedisReactionPrefixKey()).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("Del key: %s", r.getRedisReactionPrefixKey())
}

func (r *redisReactionRepository) getRedisReactionPrefixKey() string {
	if r.cfg.ServiceSettings.RedisReactionPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisReactionPrefixKey
	}

	return redisReactionPrefixKey
}
