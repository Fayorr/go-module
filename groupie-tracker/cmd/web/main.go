package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"fayokunmiosho.com/groupie-tracker/internal/models"
)

type application struct {
	logger      *slog.Logger
	artistList  []models.PageData
	artistCache map[int]models.PageData
}

func main() {
	addr := flag.String("addr", ":8000", "server port")

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	artistList, artistCache := FetchAllData()

	app := &application{
		logger:      logger,
		artistList:  artistList,
		artistCache: artistCache,
	}
	// fmt.Println("pageData", pageData)
	if artistList == nil {
		app.logger.Error("No Data fetched")
		os.Exit(1)
	}

	app.logger.Info("Starting server at port http://localhost:8000")
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		log.Fatal(err)
		return
	}
}
