package main

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./templates/index.html",
		"./templates/index.css",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w, r)
		return
	}
	err = ts.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println(err.Error())
		http.NotFound(w, r)
		return
	}
}

func DisplayArt(w http.ResponseWriter, r *http.Request) {

	sentence := r.FormValue("textInput")
	banner := r.FormValue("bannerType")

	// 1. Check for missing data
	if sentence == "" {
		http.Error(w, "400 Bad Request - Invalid Input", http.StatusBadRequest)
		return
	}

	// 2. Validation of the banner type
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" || banner == "" {
		http.Error(w, "400 Bad Request - Invalid Banner", http.StatusBadRequest)
		return
	}

	// 3. Loop through every individual character in the sentence
	for _, char := range sentence {
		// If the character is NOT a newline (10) AND NOT a carriage return (13)
		// AND it falls outside the standard 32-126 range...
		if char != '\n' && char != '\r' && (char < 32 || char > 126) {
			http.Error(w, "400 Bad Request - Non-ASCII Character Detected", http.StatusBadRequest)
			return
		}
	}
	log.Print(banner) // log banner selected
	result, err := Runner(sentence, banner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	files := []string{
		"./templates/ascii.html",
		"./templates/ascii.css",
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	data := struct {
		Result string
	}{
		Result: result,
	}

	tmpl.ExecuteTemplate(w, "ascii", data)
}
