package dao

import (
	database "StuService-Go/internal/pkg/databse"
	"context"
	"time"
)

func RedisSetKeyVal(ctx context.Context, key string, val string, expire time.Duration) error {
	return database.RedisDB.SetEX(ctx, key, val, expire).Err()
}

func RedisGetKeyVal(ctx context.Context, key string) (string, error) {
	return database.RedisDB.Get(ctx, key).Result()
}

func RedisDelKeyVal(ctx context.Context, key string) error {
	return database.RedisDB.Del(ctx, key).Err()
}
