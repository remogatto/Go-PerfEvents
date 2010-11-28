package perf

import "os"
import "syscall"

func newPerfCounterObject() (*PerfCounter, os.Error) {
	return &PerfCounter{attr: Attr{}, fd: make(map[int](*os.File))}, nil
}

func gettid() int {
	return syscall.Gettid();
}
