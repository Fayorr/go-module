package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./templates/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w, r)
	}
	err = ts.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w, r)
	}
}

func displayArt(w http.ResponseWriter, r *http.Request) {
	sentence := r.FormValue("textInput")
	banner := r.FormValue("bannerType")
	log.Print(banner)
	result := Runner(sentence, banner)

	tmpl, err := template.ParseFiles("./templates/ascii.html")

	if err != nil {
		log.Println(err.Error())
	}
	data := struct {
		Result string
	}{
		Result: result,
	}
	 
	tmpl.Execute(w, data)
}
