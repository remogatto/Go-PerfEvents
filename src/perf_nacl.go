package perf

import "os"

func newPerfCounterObject() (*PerfCounter, os.Error) {
	return nil, os.NewError("there is no support for Linux performance counters on NaCl")
}

func gettid() int {
	panic("not supported")
}
