package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/anhnmt/go-fingerprint"
)

func init() {
	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			TimeFormat: time.RFC3339,
			AddSource:  true,
		}),
	))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fingerprint.NewFingerprint(r)

		w.Write([]byte("Hello World"))
	})

	addr := ":8080"
	log.Printf("Listening on http://localhost%s\n", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		return
	}
}
