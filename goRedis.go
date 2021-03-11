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

// from https://tutorialedge.net/golang/go-redis-tutorial/
func RedisGoAddValue() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	// we can call set with a `Key` and a `Value`.
	err := client.Set(context.TODO(),"name", "Elliot", 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Redis key/value pair added.")
}

// from https://tutorialedge.net/golang/go-redis-tutorial/
func RedisGoValue() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	val, err := client.Get(context.TODO(),"name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}