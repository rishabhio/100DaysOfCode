package main

import "github.com/go-redis/redis"
import (
	"fmt"

)

func SetupClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	val, err := client.Get("test").Result()
	if err != nil {
		panic( err )
	}
	fmt.Println("test", val)
}

func main() {
	fmt.Println("Hello World of Redis!")
	SetupClient()
}