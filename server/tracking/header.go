package tracking

import (
	pirschip "github.com/pirsch-analytics/pirsch/v6/pkg/tracker/ip"
	"github.com/pirsch-analytics/tour/server/cfg"
)

// getHeaderParser returns a list of parsers to extract the visitor IP address from HTTP headers.
func getHeaderParser() []pirschip.HeaderParser {
	parser := cfg.Get().IPHeader

	if len(parser) > 0 {
		headerParser := make([]pirschip.HeaderParser, 0, len(parser))

		for _, parser := range parser {
			if parser == "Hetzner" {
				headerParser = append(headerParser, getHetznerLBHeaderParser())
				continue
			} else if parser == "Caddy" {
				headerParser = append(headerParser, getCaddyHeaderParser())
				continue
			}

			for _, p := range pirschip.DefaultHeaderParser {
				if parser == p.Header {
					headerParser = append(headerParser, p)
					break
				}
			}
		}

		return headerParser
	}

	return []pirschip.HeaderParser{} // none
}
