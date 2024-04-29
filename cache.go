package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	Set(key string, value interface{}, expiredTime int) error
	Get(key string) (interface{}, error)
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // set password if required
		DB:       0,  // use default DB
	})

	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Set(key string, value interface{}, expiredTime int) error {
	// set key-value pair with expiration time in seconds, if expiredTime is 0, it will never expire
	// convert expiredTime to duration first
	duration := time.Duration(expiredTime) * time.Second
	err := r.client.Set(key, value, duration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key-value pair: %v", err)
	}

	return nil
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	return r.client.Get(key).Result()
}
