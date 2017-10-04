package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}

func imageHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	imageURL := vars["imageUrl"]

	redisURL := getEnv("REDIS_URL", "localhost:6379")
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := client.Get(imageURL).Result()
	if err == redis.Nil {
		fmt.Println("key does not exists")
		getImage(response, request, imageURL)
		return
		// in this case grab the image, resize, save and respond
	} else if err != nil {
		panic(err)
	}

	getImage(response, request, val)
}
