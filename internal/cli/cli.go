// Package cli implements the console3 command line.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"

	"github.com/iter8ai/console3/internal/config"
)

// Run executes the command line and returns the process exit code.
func Run(args []string, getenv func(string) string, stdout, stderr io.Writer) int {
	cfg, err := config.Load(getenv)
	if err != nil {
		_, _ = fmt.Fprintln(stderr, "console3:", err)
		return 1
	}

	fs := flag.NewFlagSet("console3", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {}
	verbose := fs.Bool("verbose", cfg.Verbose, "print extra output (env CONSOLE3_VERBOSE)")

	usage := func(w io.Writer) {
		_, _ = fmt.Fprintln(w, "Usage: console3 [--verbose]")
		fs.SetOutput(w)
		fs.PrintDefaults()
		fs.SetOutput(stderr)
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			usage(stdout)
			return 0
		}
		usage(stderr)
		return 2
	}

	_, _ = fmt.Fprintln(stdout, "console3")
	if *verbose {
		_, _ = fmt.Fprintln(stdout, "verbose output enabled")
	}
	return 0
}
