package ab

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// Experiment alternates between different variants to split them evenly across visitors.
type Experiment struct {
	name     string
	variants []string
	current  int
	m        sync.Mutex
}

// New creates a new A/B testing experiment for the given name and variants.
func New(name string, experiments []string) *Experiment {
	return &Experiment{
		name:     name,
		variants: experiments,
	}
}

// Next returns the experiment name and variant for the next visitor.
// The selection is stored inside a cookie.
// The operation is thread-safe.
func (e *Experiment) Next(w http.ResponseWriter, r *http.Request) (string, string) {
	if len(e.variants) == 0 {
		return "", ""
	}

	// Check if a cookie with the experiment name exists and return the value if found.
	c, _ := r.Cookie(e.getCookieName())

	if c != nil {
		return c.Name, c.Value
	}

	// Select the next variant for this visitor and store it inside a cookie.
	e.m.Lock()
	variant := e.variants[e.current]
	e.current++

	if e.current >= len(e.variants) {
		e.current = 0
	}

	e.m.Unlock()
	http.SetCookie(w, &http.Cookie{
		Name:     e.getCookieName(),
		Value:    variant,
		Expires:  time.Now().UTC().Add(time.Hour * 24),
		Secure:   false,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
	slog.Info("A/B testing experiment", "name", e.name, "variant", variant)
	return e.name, variant
}

func (e *Experiment) getCookieName() string {
	return fmt.Sprintf("experiment-%s", e.name)
}
