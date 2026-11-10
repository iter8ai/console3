package cli

import (
	"bytes"
	"strings"
	"testing"
)

func run(args []string, envs map[string]string) (int, string, string) {
	var out, errb bytes.Buffer
	code := Run(args, func(k string) string { return envs[k] }, &out, &errb)
	return code, out.String(), errb.String()
}

func TestHelp(t *testing.T) {
	for _, flag := range []string{"--help", "-h"} {
		code, out, _ := run([]string{flag}, nil)
		if code != 0 {
			t.Errorf("%s: exit %d, want 0", flag, code)
		}
		if !strings.Contains(out, "Usage") || !strings.Contains(out, "--verbose") {
			t.Errorf("%s: stdout lacks usage: %q", flag, out)
		}
	}
}

func TestUnknownFlag(t *testing.T) {
	code, _, errOut := run([]string{"--nope"}, nil)
	if code != 2 {
		t.Errorf("exit %d, want 2", code)
	}
	if !strings.Contains(errOut, "Usage") {
		t.Errorf("stderr lacks usage: %q", errOut)
	}
}

func TestInvalidEnv(t *testing.T) {
	code, _, errOut := run(nil, map[string]string{"CONSOLE3_VERBOSE": "banana"})
	if code != 1 || !strings.Contains(errOut, "CONSOLE3_VERBOSE") {
		t.Errorf("exit %d, stderr %q; want 1 naming the variable", code, errOut)
	}
}

func TestFlagBeatsEnv(t *testing.T) {
	envs := map[string]string{"CONSOLE3_VERBOSE": "true"}
	_, out, _ := run(nil, envs)
	if !strings.Contains(out, "verbose") {
		t.Errorf("env true: want verbose output, got %q", out)
	}
	code, out, _ := run([]string{"--verbose=false"}, envs)
	if code != 0 || strings.Contains(out, "verbose") {
		t.Errorf("flag false: exit %d, out %q; want non-verbose", code, out)
	}
}
