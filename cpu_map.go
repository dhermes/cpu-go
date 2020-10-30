// +build !linux,!darwin

package cpu

import (
	"fmt"
	"runtime"
)

// mapCPUs is a default implementation in the absence of an actual assembly
// implementation. It returns and empty map.
func mapCPUs() (map[uint64]int, error) {
	err := fmt.Errorf(
		"%w; GOOS: %q, GOARCH: %q",
		ErrUnsupportedPlatform, runtime.GOOS, runtime.GOARCH,
	)
	return nil, err
}
