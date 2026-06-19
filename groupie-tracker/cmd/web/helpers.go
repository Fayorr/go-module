package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
var (
	uri = r.URL.RequestURI()
	method = r.Method
)
	app.logger.Error(err.Error(), "uri", uri, "method", method)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) getJSON(url string, target interface{}) error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(&target)
}