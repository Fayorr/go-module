package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", Home)
	mux.HandleFunc("POST /ascii-art", DisplayArt)
	log.Println("Starting Go server at port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err.Error())
}
