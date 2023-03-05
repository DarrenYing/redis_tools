package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	tools "redis_tools"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	disLock, err := tools.NewRedisLock(client, "lock resource")
	if err != nil {
		log.Fatal(err)
	}

	succ, err := disLock.TryLock(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	if succ {
		defer disLock.Unlock(context.Background())
	}

	succ, err = disLock.SpinLock(context.Background(), 5)
	if err != nil {
		log.Println(err)
		return
	}

	if succ {
		defer disLock.Unlock(context.Background())
	}
}
