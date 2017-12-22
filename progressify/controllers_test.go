package main_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func TestKeyThatDoesNotExistYet(t *testing.T) {
	response, err := http.Get("http://localhost:8081/https://www.w3schools.com/css/trolltunga.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, response.Status, "200 OK")
}

func TestNonExistingURL(t *testing.T) {
	response, err := http.Get("http://localhost:8081/https://www.w3schools.com/css/this_image_does_not_exist.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, response.Status, "404 Not Found")
}

func TestURLThatIsNotAnImage(t *testing.T) {
	response, err := http.Get("http://localhost:8081/https://czaplinski.io/")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, response.Status, "404 Not Found")
}

func TestGettingURLThatExistInDatabase(t *testing.T) {
	redisURL := getEnv("REDIS_URL", "localhost:6379")
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client.Set("https://www.w3schools.com/css/trolltunga.jpg", "https://www.w3schools.com/css/trolltunga.jpg", 0)
	defer client.Del("https://www.w3schools.com/css/trolltunga.jpg")

	response, err := http.Get("http://localhost:8081/https://www.w3schools.com/css/trolltunga.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, response.Status, "200 OK")
}
