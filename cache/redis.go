package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type redisCache struct {
	ctx               context.Context
	client            *redis.Client
	Prefix            string
	DefaultExpireTime time.Duration
}

func NewRedisCache(cli *redis.Client, pre string) Cache {
	return &redisCache{
		ctx:    context.Background(),
		client: cli,
		Prefix: pre,
	}
}

func (r *redisCache) Set(key string, val interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return errors.Wrapf(err, "marshal data err, value is %+v", val)
	}
	cacheKey, err := BuildCacheKey(r.Prefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}
	if expiration == 0 {
		expiration = DefaultExpireTime
	}
	err = r.client.Set(r.ctx, cacheKey, string(bytes), expiration).Err()
	if err != nil {
		return errors.Wrapf(err, "redis set err: %+v,", err)
	}
	return nil
}

func (r *redisCache) Get(key string, val interface{}) error {
	cacheKey, err := BuildCacheKey(r.Prefix, key)
	if err != nil {
		return errors.Wrapf(err, "build cache key err, key is %+v", key)
	}

	bytes, err := r.client.Get(r.ctx, cacheKey).Bytes()
	if err != nil && err != redis.Nil {
		return errors.Wrapf(err, "redis get err,key = %v,", cacheKey)
	}

	if string(bytes) == "" {
		return nil
	}
	err = json.Unmarshal(bytes, &val)
	if err != nil {
		return errors.Wrapf(err, "Unmarshal err,key = %v,cacheKey=%v", key, cacheKey)
	}
	return nil
}

func (r *redisCache) Del(keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	cacheKeys := make([]string, len(keys))

	for index, key := range keys {
		cacheKey, err := BuildCacheKey(r.Prefix, key)
		if err != nil {
			log.Printf("build cache key err: %+v, key is %+v", err, key)
			continue
		}
		cacheKeys[index] = cacheKey
	}

	err := r.client.Del(r.ctx, cacheKeys...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete error, keys is %+v", keys)
	}
	return nil
}
