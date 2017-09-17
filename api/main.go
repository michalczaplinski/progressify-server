package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// - allow using own redis
// - make

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: /")
}

func imageAddres(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	imageURL := vars["imageUrl"]

	fmt.Fprint(writer, "the passed URL was: ", imageURL)

	// fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: /{imageUrl}")

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/{imageUrl:.*}", imageAddres)

	// TODO: handle invalid imageUrl that is not a real url

	// TODO: handle when there is no image under particular url

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
