package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/bytedance/sonic"
	"github.com/lmittmann/tint"

	"github.com/anhnmt/go-fingerprint"
)

type Fingerprint struct {
	ID string `json:"id,omitempty"`
	*fingerprint.Fingerprint
}

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
		fg := fingerprint.NewFingerprint(r)

		newFingerprint := &Fingerprint{
			ID:          fg.ID(),
			Fingerprint: fg,
		}

		marshal, err := sonic.Marshal(newFingerprint)
		if err != nil {
			return
		}

		slog.Info("Fingerprint",
			slog.String("id", fg.ID()),
			slog.String("data", string(marshal)),
			slog.Any("headers", r.Header),
			slog.Any("remote-addr", r.RemoteAddr),
		)

		w.Header().Set("Content-Type", "application/json")
		w.Write(marshal)
	})

	addr := ":8080"
	slog.Info(fmt.Sprintf("Listening on http://localhost%s", addr))

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		slog.Error(fmt.Sprintf("Error: %s", err.Error()))
	}
}
