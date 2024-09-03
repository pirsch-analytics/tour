package tracking

import (
	pirschip "github.com/pirsch-analytics/pirsch/v6/pkg/tracker/ip"
	"net"
	"strings"
)

// getCaddyHeaderParser returns a parser for the X-Forwarded-For header provided by Caddy.
func getCaddyHeaderParser() pirschip.HeaderParser {
	return pirschip.HeaderParser{
		Header: "X-Forwarded-For",
		Parser: parseXForwardedForFirst,
	}
}

func parseXForwardedForFirst(value string) string {
	parts := strings.Split(value, ",")

	if len(parts) > 1 {
		ip := parts[0]

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
