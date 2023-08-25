package main

import (
	"os"

	"github.com/JeanBrov/fluxmc/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
