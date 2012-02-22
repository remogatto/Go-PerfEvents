package perf

import "errors"

func newCounterObject() (*Counter, error) {
	return nil, errors.New("there is no support for Linux performance counters on Windows")
}

func gettid() int {
	panic("not supported")
}
