package main

import (
	"log"
	"net/http"

	"github.com/anhnmt/go-fingerprint"
)

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
