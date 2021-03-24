package main

import (
	"context"
	"fmt"

	"encoding/json"
)

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

// from https://tutorialedge.net/golang/go-redis-tutorial/
func RedisGoJson() {
	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(context.TODO(), "id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(context.TODO(),"id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}