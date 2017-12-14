package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true).SkipClean(true)
	router.HandleFunc("/", indexController)

	router.HandleFunc("/{imageUrl:.*}", imageController)

	// TODO: handle invalid imageUrl that is not a real url
	// TODO: handle when there is no image under particular url

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
