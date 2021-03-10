package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// from https://tutorialedge.net/golang/go-redis-tutorial/
func RedisGo() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping(context.TODO()).Result()
	fmt.Println(pong, err)
}