package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getImage(response http.ResponseWriter, request *http.Request, imageURL string) {
	imageResp, err := http.Get(imageURL)
	if err != nil {
		http.NotFound(response, request)
		return
	}
	defer imageResp.Body.Close()

	response.Header().Set("Content-Length", fmt.Sprint(imageResp.ContentLength))
	response.Header().Set("Content-Type", imageResp.Header.Get("Content-Type"))

	_, err = io.Copy(response, imageResp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

}
