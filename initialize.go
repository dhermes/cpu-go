package cpu

import (
	"fmt"
	"runtime"
	"sync"
)

// cpuMap maps (non-consecutive) CPUID values to integer indices.
var cpuMap map[uint64]int

var cpuMapLock = sync.Mutex{}

// Initialize will construct the map of CPUs so that CPUIDs can be looked up
// at runtime.
//
// This function **could** be replaced by a true `init()` but for now this
// package is attempting to avoid implicit or hidden behavior.
func Initialize() error {
	cpuMapLock.Lock()
	defer cpuMapLock.Unlock()

	// Early exit if `cpuMap` is already set.
	if cpuMap != nil {
		return nil
	}

	var err error
	cpuMap, err = mapCPUs()
	if err != nil {
		return err
	}

	if len(cpuMap) != runtime.NumCPU() {
		return fmt.Errorf(
			"%w; CPUs found: %d, runtime.NumCPU(): %d",
			ErrProcessorCount, len(cpuMap), runtime.NumCPU(),
		)
	}

	return nil
}
