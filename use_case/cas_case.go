package main

import (
	"context"
	"fmt"
	tools "github.com/DarrenYing/redis_tools"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	succ, err := tools.NewTools(client).Cas(context.Background(),
		"cas_key", "old val", "new val")

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(succ)

	res, _ := client.Get(context.Background(), "cas_key").Result()
	fmt.Println(res)

	succ, err = tools.NewTools(client).Cad(context.Background(), "cas_key", "old val")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(succ)
}
