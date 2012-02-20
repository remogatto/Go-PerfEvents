package perf

import (
	"errors"
	"os"
)

func sys_perf_counter_open(attr *Attr, pid, cpu, group_fd int, flags uint64) (counter *os.File, err error) {
	return nil, errors.New("there is no support for Linux performance counters on Windows")
}
