package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/pages/artist.html",
		"./ui/static/css/artist.css",
	}

	// 1. Create a FuncMap containing your custom formatting rules
	funcMap := template.FuncMap{
		"formatLocation": func(rawLocation string) map[string]string {
			
			// Split "los_angeles-usa" into ["los_angeles", "usa"]
			parts := strings.Split(rawLocation, "-")
			city := ""
			country := ""

			if len(parts) > 0 {
				
				// Replace underscores with spaces and title case it ("los_angeles" -> "Los Angeles")
				caser := cases.Title(language.Und)
				city = caser.String((strings.ReplaceAll(parts[0], "_", " ")))
			}
			if len(parts) > 1 {
				// Title case the country ("usa" -> "Usa")
				caser := cases.Title(language.Und)
				country = caser.String((strings.ReplaceAll(parts[1], "_", " ")))
			}

			return map[string]string{
				"City":    city,
				"Country": country,
			}
		},
	}

	// 2. Instantiate template with Funcs BEFORE calling ParseFiles
	ts, err := template.New("artist.html").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	ts.ExecuteTemplate(w, "artist", app.artistCache[id])
}