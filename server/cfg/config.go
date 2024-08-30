package cfg

import (
	"encoding/json"
	"os"
)

var (
	// cfg is the global configuration variable.
	cfg Config
)

// Config is the application configuration.
type Config struct {
	Dev                      bool   `json:"dev"`
	Host                     string `json:"host"`
	ScriptIdentificationCode string `json:"script_identification_code"`
	ScriptSrc                string `json:"script_src"`
	ScriptHitEndpoint        string `json:"script_hit_endpoint"`
	ScriptEventEndpoint      string `json:"script_event_endpoint"`
	ScriptDev                bool   `json:"script_dev"`
	ClientSecret             string `json:"client_secret"`
}

// Load loads the configuration from file.
func Load(path string) error {
	content, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	return json.Unmarshal(content, &cfg)
}

// Get returns the configuration.
func Get() Config {
	return cfg
}
