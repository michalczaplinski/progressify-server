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

	// parse the image URL from the route URL
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
	if err == redis.Nil {
		fmt.Println("key does not exists")

		// get the original image
		imageResponse, err := getImage(imageURL)

		if err != nil {
			fmt.Println(err)
			http.NotFound(response, request)
			return
		}
		defer imageResponse.Body.Close()

		writeImageToResponse(response, imageResponse)

	} else if err != nil {
		fmt.Printf("There was an error in redis retrieving url: %s\n", imageURL)

		//TODO: maybe return something more informative

	} else {

		fmt.Printf("image exists: %s", val)
		// in this case grab the image, resize, save and respond

	}

}
