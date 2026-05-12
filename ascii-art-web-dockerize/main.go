package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", Home)
	mux.HandleFunc("POST /ascii-art", DisplayArt)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/"))))
	log.Println("Starting Go server at port http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err.Error())
}
