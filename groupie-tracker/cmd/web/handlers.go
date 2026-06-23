package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/index.html",
		"./ui/static/css/home.css",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	ts.ExecuteTemplate(w, "index", app.artistList)
}

func (app *application) artistInfo(w http.ResponseWriter, r *http.Request) {
	
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w,r)
		
	}

	files := []string{
		"./ui/html/pages/artist.html",
		"./ui/static/css/artist.css",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	ts.ExecuteTemplate(w, "artist", app.artistCache[id])
}