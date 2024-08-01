package main

import (
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
