// Command console3 is the console3 command line.
package main

import (
	"os"

	"github.com/iter8ai/console3/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Getenv, os.Stdout, os.Stderr))
}
