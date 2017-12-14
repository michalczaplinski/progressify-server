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

func getImage(imageURL string) (*http.Response, error) {
	imageResp, err := http.Get(imageURL)
	if err != nil {
		return nil, err
	}
	defer imageResp.Body.Close()
	return imageResp, nil
}

func writeImageToResponse(response http.ResponseWriter, imageResp *http.Response) {

	response.Header().Set("Content-Length", fmt.Sprint(imageResp.ContentLength))
	response.Header().Set("Content-Type", imageResp.Header.Get("Content-Type"))

	_, err := io.Copy(response, imageResp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

}
