package fingerprint

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/bytedance/sonic"
)

type Fingerprint struct {
	ID        string `json:"id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}

func NewFingerprint(r *http.Request) *Fingerprint {
	args := []any{
		slog.Any("headers", r.Header),
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}

	args = append(args, slog.Any("ip", ip))
	slog.Info("Fingerprint", args...)

	return &Fingerprint{}
}

func (f *Fingerprint) Byte() ([]byte, error) {
	bytes, err := sonic.Marshal(f)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (f *Fingerprint) String() (string, error) {
	bytes, err := f.Byte()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
