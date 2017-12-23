package main

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getImage(imageURL string) (*http.Response, error) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	imageResp, err := client.Get(imageURL)

	if err != nil {
		return nil, err
	}

	imageContentType := imageResp.Header.Get("Content-Type")

	if !strings.HasPrefix(imageContentType, "image") {
		return nil, &errWrongContentType{msg: "The image under the url has wrong content type"}
	}

	return imageResp, nil
}

func writeImageToResponse(writer http.ResponseWriter, imageBytes []byte) {

	_, err := io.Copy(writer, bytes.NewReader(imageBytes))
	if err != nil {
		log.Println(err)
		return
	}

}
