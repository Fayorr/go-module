package main

import (
	"html/template"
	"net/http"

	"fayokunmiosho.com/groupie-tracker/config"
	"fayokunmiosho.com/groupie-tracker/internal/models"
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

	artistsUrl, locationsUrl, datesUrl, relationsUrl := config.ARTIST_URL_API, config.LOCATION_URL_API, config.DATES_URL_API, config.RELATIONS_URL_API

	var artists []models.Artist
    var locations models.LocationsWrapper
    var dates models.DatesWrapper
    var relations models.RelationsWrapper

	// fetch and decode
	if err := app.getJSON(artistsUrl, &artists); err != nil {
		app.serverError(w,r,err)
		return
	}
	if err := app.getJSON(locationsUrl, &locations); err != nil {
		app.serverError(w,r,err)
		return
	}
	if err := app.getJSON(datesUrl, &dates); err != nil {
		app.serverError(w,r,err)
		return
	}
	if err := app.getJSON(relationsUrl, &relations); err != nil {
		app.serverError(w,r,err)
		return
	}

	var pageData []models.PageData

	for i := range artists {
		combined := models.PageData{
			Artist: artists[i],
			Locations: locations.Index[i],
			ConcertDates: dates.Index[i],
			Relations: relations.Index[i],
		}
		pageData = append(pageData, combined)
	}

	ts.ExecuteTemplate(w, "index", pageData)
}