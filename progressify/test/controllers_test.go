package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyThatDoesNotExistYet(t *testing.T) {
	client := connectToDB()
	defer cleanUp(client)

	response, err := http.Get("http://localhost:8081/https://w3schools.com/css/trolltunga.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "200 OK", response.Status)
}

func TestNonExistingURL(t *testing.T) {
	client := connectToDB()
	defer cleanUp(client)

	response, err := http.Get("http://localhost:8081/https://w3schools.com/css/this_image_does_not_exist.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "404 Not Found", response.Status)
}

func TestURLThatIsNotAnImage(t *testing.T) {
	client := connectToDB()
	defer cleanUp(client)

	response, err := http.Get("http://localhost:8081/https://czaplinski.io/")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "404 Not Found", response.Status)
}

func TestGettingURLThatExistInDatabase(t *testing.T) {
	client := connectToDB()
	defer cleanUp(client)

	err := client.Set("https://w3schools.com/css/trolltunga.jpg", "https://w3schools.com/css/trolltunga.jpg", 0).Err()
	if err != nil {
		panic(err)
	}

	response, err := http.Get("http://localhost:8081/https://w3schools.com/css/trolltunga.jpg")
	if err != nil {
		t.Errorf("getting the image failed")
	}
	assert.Equal(t, "200 OK", response.Status)
}
