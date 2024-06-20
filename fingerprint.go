package fingerprint

import (
	"net/http"

	"github.com/bytedance/sonic"
)

type Fingerprint struct {
	IpAddress *FpIpAddress `json:"ip_address,omitempty"`
}

func NewFingerprint(r *http.Request) *Fingerprint {
	fingerprint := &Fingerprint{
		IpAddress: ParseIpAddress(r),
	}

	return fingerprint
}

func (f *Fingerprint) ID() string {
	return ""
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
