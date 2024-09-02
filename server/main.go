package main

import (
	"github.com/pirsch-analytics/tour/server/ab"
	"github.com/pirsch-analytics/tour/server/cfg"
	"github.com/pirsch-analytics/tour/server/tpl"
	"log/slog"
	"net/http"
	"os"
)

var (
	experimentPhoneHeader = ab.New("phone-header", []string{"simple", "colorful-cta"})
	experimentPadHeader   = ab.New("pad-header", []string{"simple", "colorful-cta"})
	experimentWatchHeader = ab.New("watch-header", []string{"simple", "colorful-cta"})
)

// home handles requests to the home page and to all pages which might not be found.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		tpl.ExecTpl(w, r, "home.html", nil)
	} else {
		w.WriteHeader(http.StatusNotFound)
		tpl.ExecTpl(w, r, "not-found.html", nil)
	}
}

// product handles requests to the product page for a specific product.
func product(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "product.html", struct {
		Slug string
	}{
		r.PathValue("slug"),
	})
}

// checkout handles requests to the checkout page for a specific product.
func checkout(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "checkout.html", struct {
		Slug string
	}{
		r.PathValue("slug"),
	})
}

// thankYou handles requests to the thank-you page.
func thankYou(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "thank-you.html", nil)
}

// thankYouNewsletter handles requests to the thank-you newsletter page.
func thankYouNewsletter(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "newsletter-thank-you.html", nil)
}

// thankYouContact handles requests to the thank-you contact page.
func thankYouContact(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "contact-thank-you.html", nil)
}

// contact handles requests to the contact page.
func contact(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "contact.html", nil)
}

// phone handles requests to the phone landing page.
func phone(w http.ResponseWriter, r *http.Request) {
	experiment, variant := experimentPhoneHeader.Next(w, r)
	tpl.ExecTpl(w, r, "phone.html", struct {
		Experiment string
		Variant    string
	}{
		experiment,
		variant,
	})
}

// pad handles requests to the pad landing page.
func pad(w http.ResponseWriter, r *http.Request) {
	experiment, variant := experimentPadHeader.Next(w, r)
	tpl.ExecTpl(w, r, "pad.html", struct {
		Experiment string
		Variant    string
	}{
		experiment,
		variant,
	})
}

// watch handles requests to the watch landing page.
func watch(w http.ResponseWriter, r *http.Request) {
	experiment, variant := experimentWatchHeader.Next(w, r)
	tpl.ExecTpl(w, r, "watch.html", struct {
		Experiment string
		Variant    string
	}{
		experiment,
		variant,
	})
}

// main is the entry point for the application.
func main() {
	// Load the configuration.
	path := "config.json"

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	if err := cfg.Load(path); err != nil {
		slog.Error("Error loading configuration", "err", err)
		return
	}

	// Load the templates.
	if err := tpl.LoadTemplates(cfg.Get().Dev); err != nil {
		slog.Error("Error loading templates", "err", err)
		return
	}

	// Add handler functions for the server.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/product/{slug}", product)
	http.HandleFunc("/checkout/{slug}", checkout)
	http.HandleFunc("/thank-you", thankYou)
	http.HandleFunc("/newsletter-thank-you", thankYouNewsletter)
	http.HandleFunc("/contact-thank-you", thankYouContact)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/phone", phone)
	http.HandleFunc("/pad", pad)
	http.HandleFunc("/watch", watch)
	http.HandleFunc("/", home)

	// Start the server on port 8080.
	slog.Info("Starting server", "host", cfg.Get().Host)

	if err := http.ListenAndServe(cfg.Get().Host, nil); err != nil {
		slog.Error("Error starting server", "err", err)
	}
}
