package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"fayokunmiosho.com/groupie-tracker/config"
	"fayokunmiosho.com/groupie-tracker/internal/models"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
var (
	uri = r.URL.RequestURI()
	method = r.Method
)
	app.logger.Error(err.Error(), "uri", uri, "method", method)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

type InfoTypes interface {
	[]models.Artist | models.ConcertDates | models.Locations | models.Relations | models.DatesWrapper | models.LocationsWrapper | models.RelationsWrapper
}

func getJSON [T InfoTypes] (url string, target *T) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status: %d", resp.StatusCode)
	}


	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return err
	}
	
	return nil
}

func FetchAllData() ([]models.PageData, map[int]models.PageData) {
	artistsUrl, locationsUrl, datesUrl, relationsUrl := config.ARTIST_URL_API, config.LOCATION_URL_API, config.DATES_URL_API, config.RELATIONS_URL_API

	var artists []models.Artist
	var locations models.LocationsWrapper
	var dates models.DatesWrapper
	var relations models.RelationsWrapper

	// fetch and decode
	var wg sync.WaitGroup
	wg.Add(4)

	errChan := make(chan error,1) // create a channel for errors to avoid race conditions 
	//r.WithContext(context.WithValue())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	go func() {
		defer wg.Done()
		if err := getJSON(artistsUrl, &artists); err != nil {
			select {
				case <- ctx.Done():
				return
			case errChan <- err:
				fmt.Println(err)
			default:
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := getJSON(locationsUrl, &locations); err != nil {
			select {
			case  <- ctx.Done():
				return
			case errChan <- err:
				fmt.Println(err)
			default:
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := getJSON(datesUrl, &dates); err != nil {
			select {
			case <- ctx.Done():
				return
			case errChan <- err:
				fmt.Println(err)
			default:
			}
		}
	}()
	go func() {
		defer wg.Done()
		if err := getJSON(relationsUrl, &relations); err != nil {
			select {
			case <- ctx.Done():
				return
			case errChan <- err:
				fmt.Println(err)
			default:
			}
		}
	}()

	// wait
	wg.Wait()

	
	select {
	case  <- errChan:
		log.Println("Error goroutine: fetching data")
		return nil, nil
	default:
	}
	
	close(errChan) // close channel

	var pageData []models.PageData
	pageDataMap := make(map[int]models.PageData)
	
	for i := range artists {
		combined := models.PageData{
			Artist:       artists[i],
			Locations:    locations.Index[i],
			ConcertDates: dates.Index[i],
			Relations:    relations.Index[i],
		}
		pageData = append(pageData, combined)
		pageDataMap[artists[i].ID] = combined
	}

	return pageData, pageDataMap
}