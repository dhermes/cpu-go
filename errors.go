package cpu

import (
	"errors"
)

var (
	// ErrProcessorCount is the error returned when the processor count from
	// `mapCPUs()` does not match `runtime.NumCPU()`.
	ErrProcessorCount = errors.New("Processor information parsing does not match expected CPU count")
	// ErrUnsupportedPlatform indicates an OS / Architecture / Platform is not supported.
	ErrUnsupportedPlatform = errors.New("Platform not supported")
)
