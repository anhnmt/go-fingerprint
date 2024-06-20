package fingerprint

import (
	"net"
	"net/http"
	"strings"
)

const (
	XClientIp        = "X-Client-Ip"
	XForwardedFor    = "X-Forwarded-For"
	CFConnectingIp   = "CF-Connecting-Ip"
	FastlyClientIp   = "Fastly-Client-Ip"
	TrueClientIp     = "True-Client-Ip"
	XRealIp          = "X-Real-Ip"
	XClusterClientIp = "X-Cluster-Client-Ip"
	XForwarded       = "X-Forwarded"
	ForwardedFor     = "Forwarded-For"
	Forwarded        = "Forwarded"
)

// Standard headers list
var requestHeaders = []string{
	XClientIp,
	XForwardedFor,
	CFConnectingIp,
	FastlyClientIp,
	TrueClientIp,
	XRealIp,
	XClusterClientIp,
	XForwarded,
	ForwardedFor,
	Forwarded,
}

type FpIpAddress struct {
	Value string `json:"value,omitempty"`
}

func ParseIpAddress(r *http.Request) *FpIpAddress {
	clientIp := GetClientIp(r)
	if clientIp == "" {
		return nil
	}

	return &FpIpAddress{
		Value: clientIp,
	}
}

func GetClientIp(r *http.Request) string {
	if len(r.Header) > 0 {
		for _, header := range requestHeaders {
			switch header {
			case XForwardedFor: // Load-balancers (AWS ELB) or proxies.
				host, correctIP := getClientIPFromXForwardedFor(r.Header.Get(header))
				if correctIP {
					return host
				}
			default:
				if host := r.Header.Get(header); isCorrectIP(host) {
					return host
				}
			}
		}
	}

	// remote address checks.
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && isCorrectIP(host) {
		return host
	}

	return ""
}

// getClientIPFromXForwardedFor  - returns first known ip address else return empty string
func getClientIPFromXForwardedFor(headers string) (string, bool) {
	if headers == "" {
		return "", false
	}
	// x-forwarded-for may return multiple IP addresses in the format: "client IP, proxy 1 IP, proxy 2 IP"
	// Therefore, the right-most IP address is the IP address of the most recent proxy
	// and the left-most IP address is the IP address of the originating client.
	forwardedIps := strings.Split(headers, ",")
	for _, ip := range forwardedIps {
		// header can contain spaces too, strip those out.
		ip = strings.TrimSpace(ip)
		// make sure we only use this if it's ipv4 (ip:port)
		if split := strings.Split(ip, ":"); len(split) == 2 {
			ip = split[0]
		}
		if isCorrectIP(ip) {
			return ip, true
		}
	}
	return "", false
}

// isCorrectIP - return true if ip string is valid textual representation of an IP address, else returns false
func isCorrectIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
