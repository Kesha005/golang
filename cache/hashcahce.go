package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	Add  string
	Pass string
}

func (rdata *RedisStorage) SetClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     rdata.Add,
		Password: rdata.Pass,
	})
	return client, nil
}

func main() {

	var ctx = context.Background()
	rdata := RedisStorage{Add: "localhost:6379", Pass: ""}
	client, err := rdata.SetClient()
	if err != nil {
		fmt.Println("error ", err)
	}

	seterr := client.Set(ctx, "key", "value", 0).Err()

	if seterr != nil {
		fmt.Println("error caching ", seterr)
		return
	}

	val, geterr := client.Get(ctx, "key").Result()
	if geterr != nil {
		fmt.Print("error")
		return

	}

	fmt.Println("The value of key is : ", val)

}
