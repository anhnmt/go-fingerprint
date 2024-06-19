package fingerprint

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/bytedance/sonic"
)

type Fingerprint struct {
}

func NewFingerprint(r *http.Request) *Fingerprint {
	log.Printf("headers: %v\n\n", r.Header)

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("ip: %v\n\n", ip)
	log.Printf("======================================\n\n")

	return &Fingerprint{}
}

func (f *Fingerprint) JSON() (string, error) {
	json, err := sonic.Marshal(f)
	if err != nil {
		return "", fmt.Errorf("parse json failed: %v\n", err)
	}

	return string(json), nil
}
