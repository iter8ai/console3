package config

import (
	"strings"
	"testing"
)

func env(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestLoad(t *testing.T) {
	cfg, err := Load(env(nil))
	if err != nil || cfg.Verbose {
		t.Fatalf("empty env: got %+v, %v; want Verbose false", cfg, err)
	}
	cfg, err = Load(env(map[string]string{"CONSOLE3_VERBOSE": "true"}))
	if err != nil || !cfg.Verbose {
		t.Fatalf("true: got %+v, %v; want Verbose true", cfg, err)
	}
	_, err = Load(env(map[string]string{"CONSOLE3_VERBOSE": "banana"}))
	if err == nil || !strings.Contains(err.Error(), "CONSOLE3_VERBOSE") {
		t.Fatalf("banana: got %v; want error naming CONSOLE3_VERBOSE", err)
	}
}
