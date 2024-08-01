package main

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

func Test_strings(t *testing.T) {
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
			strings(tt.args.client)
		})
	}
}
