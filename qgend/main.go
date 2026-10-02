package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.InfoContext(ctx, "Listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
