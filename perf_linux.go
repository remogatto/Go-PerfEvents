package perf

import "os"
import "syscall"

func newCounterObject() (*Counter, error) {
	return &Counter{attr: Attr{}, fd: make(map[int](*os.File))}, nil
}

func gettid() int {
	return syscall.Gettid()
}
