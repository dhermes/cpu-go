package cpu

import (
	"bytes"
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

// mapCPUs invokes `sysctl machdep.cpu.core_count` on macOS to determine the
// number of logical cores, then creates an "identity map" from `i -> i` using
// 0-indexed markers for each core.
func mapCPUs() (map[uint64]int, error) {
	// NOTE: The command `sysctl machdep.cpu.core_count` gives the number of
	//       physical cores while the one used below gives the number of
	//       logical cores.
	cmd := exec.Command("sysctl", "machdep.cpu.thread_count")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	sysctlOutput := strings.TrimSpace(out.String())
	if !strings.HasPrefix(sysctlOutput, "machdep.cpu.thread_count: ") {
		return nil, errors.New("BAD")
	}

	countStr := strings.TrimPrefix(sysctlOutput, "machdep.cpu.thread_count: ")
	u, err := strconv.ParseUint(countStr, 10, 64)
	if err != nil {
		return nil, err
	}

	cpuMap := map[uint64]int{}
	var i uint64
	for i = 0; i < u; i++ {
		cpuMap[i] = int(i)
	}
	return cpuMap, nil
}
