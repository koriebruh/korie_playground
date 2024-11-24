package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRedisConn(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "bomb", "meledagh", 15*time.Minute).Err()
	assert.Nil(t, err)

	result, err := rdb.Get(ctx, "bomb").Result()
	assert.Equal(t, "meledagh", result)
}
