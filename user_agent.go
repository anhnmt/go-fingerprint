package fingerprint

import (
	"github.com/mileusna/useragent"
)

type FpUserAgent struct {
	Raw     string   `json:"raw,omitempty"`
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
		Raw: userAgent,
		Device: &Device{
			Name: parse.Device,
			Type: deviceType(parse),
		},
	}

	if parse.Name != "" {
		ua.Browser = &Browser{
			Name:    parse.Name,
			Version: parse.Version,
		}
	}

	if parse.OS != "" {
		ua.OS = &OS{
			Name:    parse.OS,
			Version: parse.OSVersion,
		}
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
