package pkg

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	rdb 	*redis.Client
	ctx 	context.Context
}

// Handle connection with Redis
func GetRedisClient() RedisClient {
	client := RedisClient{
		ctx: context.Background(),
		rdb: redis.NewClient(&redis.Options{
			Addr: 		os.Getenv("REDIS_ADDR"),
			Password:	"",
			DB:			0,
		}),
	}

	return client
}

// Check connection
func (r *RedisClient) Ping() {
	if err := r.rdb.Ping(r.ctx).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Redis connection successfully")
}


func (r *RedisClient) Set(key string, val string, exp time.Duration) {	
	if err := r.rdb.Set(r.ctx, key, val, exp).Err(); err != nil {
		panic(err)
	}
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.rdb.Get(r.ctx, key).Result()
}