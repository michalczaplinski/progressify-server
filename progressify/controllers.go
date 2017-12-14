package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func indexController(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}

func imageController(response http.ResponseWriter, request *http.Request) {

	// parse the image URL from the querystring
	vars := mux.Vars(request)
	imageURL := vars["imageUrl"]

	// connect to redis
	// TODO: move the db connection to another file
	redisURL := getEnv("REDIS_URL", "localhost:6379")
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// try to get the image from redis
	val, err := client.Get(imageURL).Result()

	// if there is no image key in redis
	if err != nil {
		fmt.Printf("There is no url: %s", imageURL)

	} else if err == redis.Nil {

		fmt.Println("key does not exists")

		// get the original image
		imageResponse, err := getImage(val)
		if err != nil {
			http.NotFound(response, request)
			return
		}

		writeImageToResponse(response, imageResponse)

	} else {
		fmt.Println("image exists")
		// in this case grab the image, resize, save and respond

	}

}
