package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		name           string // E.g., "Valid POST", "Missing Text", "Bad Banner"
		method         string
		path           string
		formData       string // The fake textInput and bannerType
		expectedStatus int    // E.g., 200, 400, or 404
	}{
		{
			name:           "Bad Input - 400",
			method:         "POST",
			path:           "/ascii-art",
			formData:       "textInput=😍&bannerType=standard",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Valid Request - 200",
			method:         "GET",
			path:           "/",
			formData:       "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Not Found - 404",
			method:         "GET",
			path:           "/random-page",
			formData:       "",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Not Request - 405",
			method:         "GET",
			path:           "/ascii-art",
			formData:       "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", Home)
	mux.HandleFunc("POST /ascii-art", DisplayArt)

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			r, err := http.NewRequest(tt.method, tt.path, strings.NewReader(tt.formData))
			r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			if err != nil {
				t.Fatal(err)
			}
			mux.ServeHTTP(rr, r)

			if rr.Code != tt.expectedStatus {
				t.Errorf("For %s, expected %d but got %d", tt.name, tt.expectedStatus, rr.Code)
			}
		})
	}

}
