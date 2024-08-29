package main

import (
	"github.com/pirsch-analytics/tour/server/tpl"
	"log/slog"
	"net/http"
	"os"
)

// home handles requests to the home page and to all pages which might not be found.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		tpl.ExecTpl(w, r, "home.html", nil)
	} else {
		tpl.ExecTpl(w, r, "not-found.html", nil)
	}
}

// product handles requests to product page for a specific product.
func product(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "product.html", struct {
		Slug string
	}{
		r.PathValue("slug"),
	})
}

// checkout handles requests to checkout page for a specific product.
func checkout(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "checkout.html", struct {
		Slug string
	}{
		r.PathValue("slug"),
	})
}

// contact handles requests to contact page.
func contact(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "contact.html", nil)
}

// main is the entry point for the application.
func main() {
	// Dev mode is used to live reload templates.
	devMode := len(os.Args) > 1 && os.Args[1] == "dev"

	// Load the templates.
	if err := tpl.LoadTemplates(devMode); err != nil {
		return
	}

	// Add handler functions for the server.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/product/{slug}", product)
	http.HandleFunc("/checkout/{slug}", checkout)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/", home)

	// Start the server on port 8080.
	slog.Info("Starting server on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Error starting server", "err", err)
	}
}
