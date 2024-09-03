package tracking

import (
	sdk "github.com/pirsch-analytics/pirsch-go-sdk/v2/pkg"
	pirschip "github.com/pirsch-analytics/pirsch/v6/pkg/tracker/ip"
	"github.com/pirsch-analytics/tour/server/cfg"
	"log/slog"
	"net/http"
	"strings"
)

var (
	client       *sdk.Client
	headerParser []pirschip.HeaderParser
)

// Init sets up server-side tracking if the client secret is configured.
func Init() {
	if cfg.Get().ClientSecret != "" {
		client = sdk.NewClient("", cfg.Get().ClientSecret, &sdk.ClientConfig{
			BaseURL: cfg.Get().BaseURL,
		})
		headerParser = getHeaderParser()
	}
}

// PageView tracks a page view. Tags are optional.
// The request is made in the backend, so that the visitor doesn't have to wait for it to complete.
func PageView(r *http.Request, tags map[string]string) {
	go func() {
		if client != nil {
			if err := client.PageView(r, &sdk.PageViewOptions{
				IP:   pirschip.Get(r, headerParser, nil),
				Tags: tags,
			}); err != nil {
				slog.Error("Error tracking page view", "err", err)
			}
		}
	}()
}

// Event tracks an event. Metadata and tags are optional.
// The request is made in the backend, so that the visitor doesn't have to wait for it to complete.
func Event(r *http.Request, name string, meta map[string]string, tags map[string]string) {
	go func() {
		if client != nil {
			if err := client.Event(name, 0, meta, r, &sdk.PageViewOptions{
				IP:   pirschip.Get(r, headerParser, nil),
				Tags: tags,
			}); err != nil {
				slog.Error("Error tracking event", "err", err)
			}
		}
	}()
}

// EventFromURL returns a middleware tracking events using special URL parameters.
// These can be p_event for the event name, p_meta_<key> for metadata, p_path for the path the event was triggered on.
// p_event must be set.
// Plus characters will be replaced with a slash for the values, except for the path.
func EventFromURL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go func() {
			if client != nil {
				// Get the event name.
				q := r.URL.Query()
				name := q.Get("p_event")

				if name != "" {
					name = strings.ReplaceAll(name, "+", " ")

					// Manipulate the original path if p_path is set.
					path := q.Get("p_path")

					if path != "" {
						r.URL.Path = path
					}

					// Extract all metadata fields.
					meta := make(map[string]string)

					for key, values := range q {
						if strings.HasPrefix(key, "p_meta_") && len(values) > 0 {
							meta[key[len("p_meta_"):]] = values[0]
						}
					}

					// Send the event.
					if err := client.Event(name, 0, meta, r, &sdk.PageViewOptions{
						IP: pirschip.Get(r, headerParser, nil),
					}); err != nil {
						slog.Error("Error tracking event from URL", "err", err)
					}
				}
			}
		}()

		next(w, r)
	}
}
