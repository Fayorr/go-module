package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", Home)
	mux.HandleFunc("POST /ascii-art", DisplayArt)
	fmt.Println("Starting Go Server at PORT:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
