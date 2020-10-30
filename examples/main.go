package main

import (
	"fmt"
	"os"

	cpu "github.com/dhermes/cpu-go"
)

func run() error {
	return cpu.Initialize()
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
