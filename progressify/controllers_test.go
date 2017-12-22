package main_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func setup() *redis.Client {
	redisURL := getEnv("REDIS_URL", "localhost:6379")
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.FlushDB().Err()
	if err != nil {
		panic(err)
	}
	return client
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func TestKeyThatDoesNotExistYet(t *testing.T) {
	setup()
	response, err := http.Get("http://localhost:8081/https://www.w3schools.com/css/trolltunga.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "200 OK", response.Status)
}

func TestNonExistingURL(t *testing.T) {
	setup()
	response, err := http.Get("http://localhost:8081/https://w3schools.com/css/this_image_does_not_exist.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "404 Not Found", response.Status)
}

func TestURLThatIsNotAnImage(t *testing.T) {
	setup()
	response, err := http.Get("http://localhost:8081/https://czaplinski.io/")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "404 Not Found", response.Status)
}

func TestGettingURLThatExistInDatabase(t *testing.T) {
	client := setup()

	err := client.Set("https://www.w3schools.com/css/trolltunga.jpg", "https://www.w3schools.com/css/trolltunga.jpg", 0).Err()
	if err != nil {
		panic(err)
	}

	response, err := http.Get("http://localhost:8081/https://w3schools.com/css/trolltunga.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "200 OK", response.Status)
}
