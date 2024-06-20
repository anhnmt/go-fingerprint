package fingerprint

import (
	"net/http"

	"github.com/bytedance/sonic"
	murmur3 "github.com/yihleego/murmurhash3"
)

type HashFunc func(data []byte) string

type Fingerprint struct {
	hashFunc HashFunc

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

	return fingerprint
}

func (f *Fingerprint) ID() string {
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
