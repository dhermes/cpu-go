// +build !linux

package cpu

// mapCPUs is a default implementation in the absence of an actual assembly
// implementation. It returns and empty map.
func mapCPUs() (map[uint64]int, error) {
	return map[uint64]int{}, nil
}
