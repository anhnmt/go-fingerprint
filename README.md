# go-fingerprint

**Server-side fingerprint library for Golang `net/http`**

## Features
- Generate fingerprint for each browser or device

## Installation

``` 
go get -u github.com/anhnmt/go-fingerprint
```

## Usage
```go
package main

import (
    "fmt"
    "log/slog"
    "net/http"

    "github.com/anhnmt/go-fingerprint"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fg := fingerprint.NewFingerprint(r)

        marshal, err := fg.Bytes()
        if err != nil {
            return
        }

        slog.Info("Fingerprint",
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

```

## Fingerprint example: 
```json
{
    "id": "25c4164bea78a14ffedc40fe120ebecb",
    "ip_address": {
        "value": "xxx.xxx.xxx.xxx"
    },
    "user_agent": {
        "browser": {
            "name": "Chrome",
            "version": "xxx"
        },
        "os": {
            "name": "Windows",
            "version": "10.0"
        },
        "device": {
            "type": "Desktop"
        }
    }
}
```

## License

MIT © [anhnmt](https://github.com/anhnmt)