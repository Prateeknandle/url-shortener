package main

import (
	"log"
	"net/http"

	"github.com/Prateeknandle/url-shortener/apis"
)

func main() {
	//registering routes
	router := apis.NewRouter()

	//connecting to the port
	log.Println("server is listening on port : 3000")
	log.Fatal(http.ListenAndServe("3000", router))
}
