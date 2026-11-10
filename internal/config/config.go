// Package config reads console3's configuration from the environment.
package config

import (
	"fmt"
	"strconv"
)

// Config holds every value console3 reads from the environment at start-up.
type Config struct {
	Verbose bool
}

// Load reads the configuration through getenv.
func Load(getenv func(string) string) (Config, error) {
	var cfg Config
	if v := getenv("CONSOLE3_VERBOSE"); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return Config{}, fmt.Errorf("CONSOLE3_VERBOSE: %w", err)
		}
		cfg.Verbose = b
	}
	return cfg, nil
}
