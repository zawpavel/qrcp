package main

import (
	"os"

	"github.com/zawpavel/qrcp/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
