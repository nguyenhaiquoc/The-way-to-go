package main

import (
	"context"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
)

func Test_hash(t *testing.T) {
	type args struct {
		client *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test strings",
			args: args{
				client: redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash(tt.args.client)
		})
	}
}

func Test_hash_then_delete(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	tests := []struct {
		name        string
		redisClient *redis.Client
	}{
		{
			name:        "Test strings",
			redisClient: client,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userMap := map[string]string{
				"name":  "John",
				"email": "email",
				"city":  "city",
			}
			storedMap := hash_then_delete(tt.redisClient, "user:1", userMap)
			if storedMap["name"] != "John" {
				t.Errorf("Expected name to be John, got %s", storedMap["name"])
			}
			if storedMap["email"] != "email" {
				t.Errorf("Expected email to be email, got %s", storedMap["email"])
			}

			// check if the city key was deleted or not
			if _, ok := storedMap["city"]; ok {
				t.Errorf("Expected city to be deleted, got %s", storedMap["city"])
			}

		})
	}
}

func Test_scan_play(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// remove all keys

	userMap := map[string]string{
		"name":  "John",
		"email": "email value",
		"City":  "city value", // Will be ingored by the scan because tag is not matching
		"extra": "extra value",
	}
	ctx := context.Background()
	// delete user:1 key
	_, err := client.Del(ctx, "user:1").Result()
	if err != nil {
		log.Fatal("del failed", err)
	}

	// set the hash to redis
	_, err = client.HSet(ctx, "user:1", userMap).Result()
	if err != nil {
		log.Fatal("hset failed", err)
	}

	// get the hash from redis
	kvs, err := client.HGetAll(ctx, "user:1").Result()
	if err != nil {
		log.Fatal("hgetall failed", err)
	}
	// log kvs
	log.Println("kvs", kvs)
	// scan the hash to a user struct
	var user User
	err = client.HGetAll(ctx, "user:1").Scan(&user)
	if err != nil {
		log.Fatal("scan error")
	}
	log.Println("user info", user)
}
