package memory

import (
    "context"
    "github.com/redis/go-redis/v9"
)

type RedisLayer struct {
    Client *redis.Client
}

func NewRedisLayer() *RedisLayer {
    rdb := redis.NewClient(&redis.Options{Addr: "redis:6379"})
    return &RedisLayer{Client: rdb}
}

func (r *RedisLayer) Put(ctx context.Context, key string, value string) error {
    return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *RedisLayer) Get(ctx context.Context, key string) (string, error) {
    return r.Client.Get(ctx, key).Result()
}
