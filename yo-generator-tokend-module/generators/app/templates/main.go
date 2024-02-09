package main

import (
	"os"

  "<%= packageName %>/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
