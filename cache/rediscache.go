package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password for local development
		DB:       0,                // Default DB
	})

	// Ping the Redis server to check the connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	seterr := client.Set(ctx,"key1","Kesha",0).Err()
	if seterr!=nil{
		fmt.Printf("Error occured during cache %s\n", seterr)
		return 
	}
	fmt.Println("Cached")


	val ,geterr := client.Get(ctx,"key1").Result()
	if geterr!=nil{
		fmt.Println("error ",geterr)
		return 
	}

	fmt.Print("Cache value is ", val)
}



