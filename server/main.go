package main

import (
	"github.com/pirsch-analytics/tour/server/tpl"
	"log/slog"
	"net/http"
	"os"
)

// home handles requests to the home page.
func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecTpl(w, r, "home.html", nil)
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
	http.HandleFunc("/", home)

	// Start the server on port 8080.
	slog.Info("Starting server on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Error starting server", "err", err)
	}
}
