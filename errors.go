package cpu

import (
	"errors"
)

var (
	// ErrProcessorCount is the error returned when the processor count from
	// `mapCPUs()` does not match `runtime.NumCPU()`.
	ErrProcessorCount = errors.New("Processor information parsing does not match expected CPU count")
)
