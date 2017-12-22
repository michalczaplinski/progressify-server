package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func indexController(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
}

func imageController(writer http.ResponseWriter, request *http.Request) {

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
			log.Println(fmt.Sprintf("fetching the image from %s failed!", imageURL))
			http.NotFound(writer, request)
			return
		}
		defer imageResponse.Body.Close()

		imageBytes, err := ioutil.ReadAll(imageResponse.Body)
		if err != nil {
			log.Println("could not read the imageResponse body!")
		}
		writeImageToResponse(writer, imageBytes)
		writer.Header().Set("Content-Length", fmt.Sprint(imageResponse.ContentLength))
		writer.Header().Set("Content-Type", imageResponse.Header.Get("Content-Type"))

		go func() {
			filename, err := saveFile(imageBytes)
			if err != nil {
				log.Println(err)
				return
			}

			err = client.Set(imageURL, filename, 0).Err()
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("Saved %s: %s", imageURL, filename)
		}()

	} else if err != nil {
		fmt.Printf("There was an error in redis retrieving url: %s\n", imageURL)

		//TODO: maybe return something more informative

	} else {
		fmt.Printf("image exists: %s\n", val)
		// in this case grab the image and respond

	}

}
