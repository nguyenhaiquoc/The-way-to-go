package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal("ping failed. could not connect", err)
	}

	strings(client)
}

func strings(client *redis.Client) {

	ctx := context.Background()

	_, err := client.Set(ctx, "foo", "bar", 0).Result()

	if err != nil {
		log.Fatal("set failed", err)
	}
	log.Println("set value successful")

	res, _ := client.Get(ctx, "foo").Result()
	log.Println("foo=", res)

	client.Set(ctx, "foo1", "bar1", 2*time.Second).Result()
	time.Sleep(3 * time.Second)

	_, err = client.Get(ctx, "foo1").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Println("foo1 expired")
		}
	}

	err = client.Incr(ctx, "uid").Err()
	if err != nil {
		log.Fatal("incr failed", err)
	}

	client.IncrBy(ctx, "uid", 3)
	uid, _ := client.Get(ctx, "uid").Result()
	log.Println("uid=", uid)

	//set with get - atomic

	client.SetArgs(ctx, "uid", 0, redis.SetArgs{Get: true})
	log.Println("uid=", client.Get(ctx, "uid").Val())

	//getrange
	client.Set(ctx, "long_string", "welcome to educative", 0)
	subString, _ := client.GetRange(ctx, "long_string", 0, 6).Result()
	log.Println("subString=", subString)
}
