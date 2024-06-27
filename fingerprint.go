package fingerprint

import (
	"context"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	murmur3 "github.com/yihleego/murmurhash3"
)

type HashFunc func(data []byte) string

type Fingerprint struct {
	hashFunc HashFunc

	ID        string       `json:"id,omitempty"`
	IpAddress *FpIpAddress `json:"ip_address,omitempty"`
	UserAgent *FpUserAgent `json:"user_agent,omitempty"`
}

func NewFingerprint(r *http.Request) *Fingerprint {
	fingerprint := &Fingerprint{
		hashFunc: func(data []byte) string {
			murmur := murmur3.New128()
			return murmur.HashBytes(data).String()
		},
		IpAddress: ParseIpAddress(r),
		UserAgent: ParseUserAgent(r.UserAgent()),
	}

	fingerprint.ID = fingerprint.GetID()
	return fingerprint
}

func NewFingerprintContext(ctx context.Context) *Fingerprint {
	fingerprint := &Fingerprint{
		hashFunc: func(data []byte) string {
			murmur := murmur3.New128()
			return murmur.HashBytes(data).String()
		},
		IpAddress: ParseIpAddressContext(ctx),
		UserAgent: ParseUserAgent(metadata.ExtractIncoming(ctx).Get("user-agent")),
	}

	fingerprint.ID = fingerprint.GetID()
	return fingerprint
}

func (f *Fingerprint) GetID() string {
	data, err := f.Bytes()
	if len(data) == 0 || err != nil {
		return ""
	}

	return f.hashFunc(data)
}

func (f *Fingerprint) Bytes() ([]byte, error) {
	bytes, err := sonic.Marshal(f)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (f *Fingerprint) String() (string, error) {
	bytes, err := f.Bytes()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (f *Fingerprint) SetHashFunc(fn HashFunc) {
	f.hashFunc = fn
}
