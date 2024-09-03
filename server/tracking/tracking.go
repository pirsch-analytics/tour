package tracking

import (
	sdk "github.com/pirsch-analytics/pirsch-go-sdk/v2/pkg"
	pirschip "github.com/pirsch-analytics/pirsch/v6/pkg/tracker/ip"
	"github.com/pirsch-analytics/tour/server/cfg"
	"log/slog"
	"net/http"
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

// Event tracks an event. Meta data and tags are optional.
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
