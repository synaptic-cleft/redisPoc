package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

type Album struct {
	Title  string
	Artist string
	Price  float64
	Likes  int
}

func RedigoSave() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Save using Redigo!")
}