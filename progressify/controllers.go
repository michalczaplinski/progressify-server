package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
			return // TODO: HTTP Response code
		}
		defer imageResponse.Body.Close()

		imageBytes, err := ioutil.ReadAll(imageResponse.Body)
		if err != nil {
			log.Println("could not read the imageResponse body!")
		}

		writeImageToResponse(writer, imageBytes)

		contentType := imageResponse.Header.Get("Content-Type")
		contentLength := fmt.Sprint(imageResponse.ContentLength)
		writer.Header().Set("Content-Length", contentLength)
		writer.Header().Set("Content-Type", contentType)

		splitFilename := strings.Split(imageURL, ".")
		extension := splitFilename[len(splitFilename)-1]
		if extension == imageURL || extension == "" {

			// TODO:
			// This will fail if the URL does not end with the image extension
			// Parse the contentType header to figure out the extension
			log.Fatal("Could not figure out the file's extension")
		}

		go func() {
			filename, err := saveFile(imageBytes, extension)
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
		// in this case grab the image and respond
		fmt.Printf("image exists: %s\n", val)
		imageBytes, err := ioutil.ReadFile(fmt.Sprintf("/tmp/%s", val))
		if err != nil {
			log.Print("There was an error reading the file from disk")
			return // TODO: HTTP Response code
		}

		//
		writeImageToResponse(writer, imageBytes)
		// writer.Header().Set("Content-Type", imageResponse.Header.Get("Content-Type"))

	}

}
