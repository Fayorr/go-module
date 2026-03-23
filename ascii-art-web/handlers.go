package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	 
	w.Header().Set("Allow", http.MethodGet)

	files := []string{
		"./templates/index.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w,r)
	}
	err = ts.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w,r)
	}
}

func displayArt(w http.ResponseWriter, r *http.Request) {

}