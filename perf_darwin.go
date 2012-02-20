package perf

import "errors"

func newPerfCounterObject() (*PerfCounter, error) {
	return nil, errors.New("there is no support for Linux performance counters on Darwin")
}

func gettid() int {
	panic("not supported")
}
