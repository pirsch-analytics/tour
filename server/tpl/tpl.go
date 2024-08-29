package tpl

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	// tpl holds all template files so that they can later be executed in the handlers.
	tpl *template.Template

	// devMode is used to live reload template files.
	devMode bool

	// m is a mutex ensuring we can concurrently access templates if in dev mode.
	m sync.RWMutex
)

// LoadTemplates loads the template files from the templates directory.
func LoadTemplates(liveReload bool) error {
	m.Lock()
	defer m.Unlock()
	tpl = template.New("").Funcs(template.FuncMap{
		"dict": dict,
	})

	if err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".html" {
			if _, e := tpl.ParseFiles(path); e != nil {
				slog.Error("Error loading template", "err", e, "path", path)
				return e
			}
		}

		return err
	}); err != nil {
		return err
	}

	devMode = liveReload
	return nil
}

// ExecTpl is a simple helper function logging eventual errors while executing a template and reloading files in development mode.
func ExecTpl(w http.ResponseWriter, r *http.Request, name string, data any) {
	if devMode {
		if err := LoadTemplates(devMode); err != nil {
			slog.Error("Error loading templates", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	m.RLock()
	defer m.RUnlock()

	if err := tpl.ExecuteTemplate(w, name, data); err != nil {
		slog.Error("Error executing template", "err", err, "name", name, "path", r.URL.Path)
	}
}

// dict is a helper function to pass key-value pairs in templates.
func dict(v ...any) map[string]any {
	dict := map[string]any{}
	lenv := len(v)

	for i := 0; i < lenv; i += 2 {
		key := strVal(v[i])

		if i+1 >= lenv {
			dict[key] = ""
			continue
		}

		dict[key] = v[i+1]
	}

	return dict
}

// strVal converts the given variable to a string.
func strVal(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}
