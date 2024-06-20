package fingerprint

import (
	"github.com/mileusna/useragent"
)

type FpUserAgent struct {
	Browser *Browser `json:"browser,omitempty"`
	OS      *OS      `json:"os,omitempty"`
	Device  *Device  `json:"device,omitempty"`
}

type Browser struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type OS struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type Device struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

func ParseUserAgent(userAgent string) *FpUserAgent {
	if userAgent == "" {
		return nil
	}

	parse := useragent.Parse(userAgent)

	ua := &FpUserAgent{
		Browser: &Browser{
			Name:    parse.Name,
			Version: parse.Version,
		},
		OS: &OS{
			Name:    parse.OS,
			Version: parse.OSVersion,
		},
		Device: &Device{
			Name: parse.Device,
			Type: deviceType(parse),
		},
	}

	return ua
}

func deviceType(ua useragent.UserAgent) string {
	if ua.Mobile {
		return "Mobile"
	}
	if ua.Tablet {
		return "Tablet"
	}
	if ua.Desktop {
		return "Desktop"
	}
	if ua.Bot {
		return "Bot"
	}

	return "Unknown"
}
