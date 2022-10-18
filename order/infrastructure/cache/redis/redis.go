package redis

import (
	"context"
	"time"
)

type (
	RedisInterface interface {
		Set(ctx context.Context, k string, x interface{}, d time.Duration) error
		Get(ctx context.Context, k string) (string, error)
		Delete(ctx context.Context, keys ...string) error
		GetKeyOfStructValue(ctx context.Context, key string, value interface{}) error
		SetKeyWithStructValue(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		IsEmptyData(e error) bool
		// Reset resets the storage and delete all keys.
		//FlushAll(ctx context.Context) error
		Init() error
	}
)
