package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis client Ping", pong)

	if err := client.Set("key", "value", 0).Err(); err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key2 does not exist")
		} else {
			panic(err)
		}
	}
	fmt.Println("key2", val2)
}
