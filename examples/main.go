package main

import (
	"fmt"
	"os"
	"sort"

	cpu "github.com/dhermes/cpu-go"
)

func run() error {
	err := cpu.Initialize()
	if err != nil {
		return err
	}

	cm, err := cpu.CurrentMap()
	if err != nil {
		return err
	}

	keys := make([]uint64, 0, len(cm))
	for k := range cm {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	fmt.Println("Current map of CPUs:")
	for _, k := range keys {
		v := cm[k]
		fmt.Printf("| %d -> %d\n", k, v)
	}

	here := cpu.Current()
	fmt.Printf("Current CPU: %d\n", here)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
