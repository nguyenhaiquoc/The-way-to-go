package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func hash(client *redis.Client) {

	ctx := context.Background()

	//set
	_, err := client.HSet(ctx, "user:1", "name", "user1", "email", "user1@foo.com").Result()
	if err != nil {
		log.Fatal("hset failed", err)
	}

	_, err = client.HSet(ctx, "user:2", []string{"name", "user2", "email", "user2@foo.com"}).Result()
	if err != nil {
		log.Fatal("hset failed", err)
	}

	_, err = client.HSet(ctx, "user:3", map[string]string{"name": "user3", "email": "user3@foo.com", "city": "New York"}).Result()
	if err != nil {
		log.Fatal("hset failed", err)
	}

	//get
	name, err := client.HGet(ctx, "user:1", "name").Result()
	if err != nil {
		log.Fatal("hget failed", err)
	}
	log.Println("name=", name)

	vals, err := client.HMGet(ctx, "user:2", "name", "email").Result()
	if err != nil {
		log.Fatal("hmget failed", err)
	}

	for _, val := range vals {
		log.Println("value=", val)
	}

	kvs, err := client.HGetAll(ctx, "user:3").Result()
	if err != nil {
		log.Fatal("hgetall failed", err)
	}

	for k, v := range kvs {
		log.Println(k, "=", v)
	}

	//structs
	var user User
	err = client.HGetAll(ctx, "user:3").Scan(&user)
	if err != nil {
		log.Fatal("scan error")
	}
	log.Println("user info", user)

	keys := client.HKeys(ctx, "user:1").Val()
	log.Println("keys -", keys)

}

type User struct {
	Email string `redis:"email"`
	Name  string `redis:"name"`
	City  string `redis:"city"`
}
