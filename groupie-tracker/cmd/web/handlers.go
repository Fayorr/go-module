package main

import (
	"html/template"
	"net/http"
)

func (app * application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/index.html",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w,r,err)
		return
	}
	ts.ExecuteTemplate(w, "index", nil)
}