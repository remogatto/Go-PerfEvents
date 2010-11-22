package perf

import "os"

func sys_perf_counter_open(attr *Attr, pid, cpu, group_fd int, flags uint64) (counter *os.File, err os.Error) {
	return nil, os.NewError("there is no support for Linux performance counters on Darwin")
}
