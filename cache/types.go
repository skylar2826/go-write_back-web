package cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (any, error)
	Delete(ctx context.Context, key string) error
}

var _ Cache = &BuildInMemoryCache{}
var _ Cache = &MaxMemoryCache{}
var _ Cache = &RedisCache{}
