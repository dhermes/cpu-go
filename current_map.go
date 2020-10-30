package cpu

import "errors"

// CurrentMap returns **a copy** of the current map of CPUs.
func CurrentMap() (map[uint64]int, error) {
	cpuMapLock.Lock()
	defer cpuMapLock.Unlock()

	if cpuMap == nil {
		return nil, errors.New("YUCK")
	}

	cpuMapCopy := map[uint64]int{}
	for k, v := range cpuMap {
		cpuMapCopy[k] = v
	}

	return cpuMapCopy, nil
}
