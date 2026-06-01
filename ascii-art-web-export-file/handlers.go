package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

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
	}
}

func displayArt(w http.ResponseWriter, r *http.Request) {

	sentence := r.FormValue("textInput")
	banner := r.FormValue("bannerType")
	// fmt.Printf("Raw received string: %q\n", sentence)
	// 1. Check for missing data
	if sentence == "" {
		http.Error(w, "400 Bad Request - Invalid Input", http.StatusBadRequest)
	}

	// 2. Validationm of the banner type
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
	log.Print(banner)
	result := Runner(sentence, banner)

	tmpl, err := template.ParseFiles("./templates/ascii.html")

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

	tmpl.Execute(w, data)
}

func downloadArt(w http.ResponseWriter, r *http.Request) {

	sentence := r.FormValue("textInput")
	banner := r.FormValue("bannerType")
	// 1. Check for missing data
	if sentence == "" {
		http.Error(w, "400 Bad Request - Invalid Input", http.StatusBadRequest)
	}

	// 2. Validationm of the banner type
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
	log.Print(banner)
	result := Runner(sentence, banner)

	


	fileName := "art.txt"
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))

	w.Write([]byte(result))
}
