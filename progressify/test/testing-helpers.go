package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func connectToDB() *redis.Client {
	var client *redis.Client
	redisURL := getEnv("REDIS_URL", "localhost:6379")

	if client != nil {
		return client
	}
	return func() *redis.Client {
		client := redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		return client
	}()
}

func cleanUp(client *redis.Client) {
	fmt.Println("cleaned up the tests")
	err := client.FlushDB().Err()
	if err != nil {
		panic(err)
	}

}
