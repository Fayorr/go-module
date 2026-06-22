package main

import (
	"html/template"
	"net/http"
	"sync"

	"fayokunmiosho.com/groupie-tracker/config"
	"fayokunmiosho.com/groupie-tracker/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/index.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	artistsUrl, locationsUrl, datesUrl, relationsUrl := config.ARTIST_URL_API, config.LOCATION_URL_API, config.DATES_URL_API, config.RELATIONS_URL_API

	var artists []models.Artist
	var locations models.LocationsWrapper
	var dates models.DatesWrapper
	var relations models.RelationsWrapper

	// fetch and decode
	var wg sync.WaitGroup
	wg.Add(4)

	errChan := make(chan error, 4) // create a channel for errors to avoid race conditions 


	go func() {
		defer wg.Done()
		if err := app.getJSON(artistsUrl, &artists); err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer wg.Done()
		if err := app.getJSON(locationsUrl, &locations); err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer wg.Done()
		if err := app.getJSON(datesUrl, &dates); err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer wg.Done()
		if err := app.getJSON(relationsUrl, &relations); err != nil {
			errChan <- err
		}
	}()

	// wait
	wg.Wait()

	close(errChan) // close channel

	for err := range errChan {
		if err != nil {
			app.serverError(w,r,err)
			return
		}
	}

	var pageData []models.PageData

	for i := range artists {
		combined := models.PageData{
			Artist:       artists[i],
			Locations:    locations.Index[i],
			ConcertDates: dates.Index[i],
			Relations:    relations.Index[i],
		}
		pageData = append(pageData, combined)
	}

	ts.ExecuteTemplate(w, "index", pageData)
}
