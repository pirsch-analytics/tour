package tracking

import (
	pirschip "github.com/pirsch-analytics/pirsch/v6/pkg/tracker/ip"
	"net"
	"strings"
)

// getHetznerLBHeaderParser returns a parser for the X-Forwarded-For header provided by the Hetzner load balancer.
func getHetznerLBHeaderParser() pirschip.HeaderParser {
	return pirschip.HeaderParser{
		Header: "X-Forwarded-For",
		Parser: parseXForwardedFor,
	}
}

func parseXForwardedFor(value string) string {
	parts := strings.Split(value, ",")

	if len(parts) > 1 {
		ip := parts[len(parts)-2]

		if strings.Contains(ip, ":") {
			host, _, err := net.SplitHostPort(ip)

			if err != nil {
				return ip
			}

			return strings.TrimSpace(host)
		}

		return strings.TrimSpace(ip)
	}

	return ""
}
