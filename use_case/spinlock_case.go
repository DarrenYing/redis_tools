package main

import (
	"context"
	tools "github.com/DarrenYing/redis_tools"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()

	ctx := context.Background()
	key := "my_lock"
	retry := 3

	disLock, _ := tools.NewRedisLock(client, key)

	wg := sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			locked, err := disLock.SpinLock(ctx, retry)
			if err != nil {
				log.Printf("goroutine %d failed to get lock: %v", id, err)
				return
			}
			if !locked {
				log.Printf("goroutine %d failed to get lock after %d retries", id, retry)
				wg.Done()
				return
			}

			log.Printf("goroutine %d acquired the lock", id)
			time.Sleep(time.Second * 2)
			_, err = disLock.Unlock(ctx)
			if err != nil {
				log.Printf("goroutine %d failed to release lock: %v", id, err)
			} else {
				log.Printf("goroutine %d released the lock", id)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
