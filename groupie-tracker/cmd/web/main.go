package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}
func main() {
	addr := flag.String("addr", ":8000", "server port")

logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

app := &application{
	logger: logger,
}
	app.logger.Info("Starting server at port http://localhost:8000")
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		log.Fatal(err)
		return
	}
}
