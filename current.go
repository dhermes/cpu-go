// +build !amd64

package cpu

// Current returns a unique identifier for the core the current goroutine is
// executing on. This function is platform dependent, and is implemented in
// `cpu_*.s`.
func Current() uint64 {
	return 0
}
