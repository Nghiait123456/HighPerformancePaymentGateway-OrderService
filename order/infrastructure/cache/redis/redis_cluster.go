package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type (
	RedisCluster struct {
		cluster *redis.ClusterClient
	}

	RedisClusterInterface interface {
		RedisInterface
	}

	Config = redis.ClusterOptions
)

func (rc *RedisCluster) Set(ctx context.Context, k string, x interface{}, d time.Duration) error {
	return rc.cluster.Set(ctx, k, x, d).Err()
}

func (rc *RedisCluster) Get(ctx context.Context, k string) (string, error) {
	return rc.cluster.Get(ctx, k).Result()
}

func (rc *RedisCluster) Delete(ctx context.Context, keys ...string) error {
	return rc.cluster.Del(ctx, keys...).Err()
}

//// Reset resets the storage and delete all keys.
//func (rc *RedisCluster) FlushAll(ctx context.Context) error {
//	return rc.cluster.FlushAll(ctx).Err()
//}

//GetKeyOfStructValue
func (rc *RedisCluster) GetKeyOfStructValue(ctx context.Context, key string, value interface{}) error {
	val, err := rc.cluster.Get(ctx, key).Result()
	if rc.IsEmptyData(err) {
		return nil
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), &value)
	if err != nil {
		return err
	}

	return nil
}

//SetKeyWithStructValue
func (rc *RedisCluster) SetKeyWithStructValue(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = rc.cluster.Set(ctx, key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc RedisCluster) IsEmptyData(e error) bool {
	return e == redis.Nil
}

func (rc *RedisCluster) Init() error {
	if err := rc.cluster.Ping(context.Background()).Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}

	return nil
}

func NewRedisCluster(cf Config) RedisClusterInterface {
	rc := RedisCluster{}
	rc.cluster = redis.NewClusterClient(&cf)
	rc.Init()

	return &rc
}
