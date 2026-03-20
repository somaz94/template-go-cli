package main

import (
	"os"

	"github.com/YOUR_USERNAME/YOUR_PROJECT/cmd/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
