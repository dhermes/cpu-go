package cpu

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// mapCPUs parses `/proc/cpuinfo` to determine a mapping from an APIC ID to a
// CPU ID.
//
// See:
// - http://en.wikipedia.org/wiki/Advanced_Programmable_Interrupt_Controller (APIC)
// - http://en.wikipedia.org/wiki/CPUID
func mapCPUs() (map[uint64]int, error) {
	cpuInfo, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return nil, err
	}

	cpuMap := map[uint64]int{}
	var processorNum int
	var apic uint64
	lines := strings.Split(string(cpuInfo), "\n")
	for i, line := range lines {
		if len(line) == 0 && i != 0 {
			cpuMap[apic] = processorNum
			processorNum = 0
			apic = 0
			continue
		}

		fields := strings.Fields(line)

		switch fields[0] {
		case "processor":
			processorNum, err = strconv.Atoi(fields[2])
		case "apicid":
			apic, err = strconv.ParseUint(fields[2], 10, 64)
		}

		if err != nil {
			return nil, err
		}
	}

	return cpuMap, nil
}
