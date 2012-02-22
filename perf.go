/*
A Go interface to performance events available in recent Linux kernels.
*/
package perf

import (
	"errors"
	"os"
	"unsafe"
)

type Counter struct {
	attr Attr
	fd   map[int](*os.File) // File descriptors for each OS thread, initially empty
}

func (attr *Attr) init(eventType uint32, event, flags uint64) {
	attr.Type = eventType
	attr.Size = _ATTR_SIZE
	attr.Config = event
	attr.Flags = flags
}

func (attr *Attr) open(pid int) (counter *os.File, err error) {
	return sys_perf_counter_open(attr, /*pid*/ pid, /*cpu*/ -1, /*group_fd*/ -1, /*flags*/ 0)
}

// Returns a new performance counter, or nil and an error.
// An error can be returned if the OS is not Linux.
func NewCounter(eventType uint32, event, flags uint64) (*Counter, error) {
	counter, err := newCounterObject()
	if err != nil {
		return nil, err
	}

	counter.attr.init(eventType, event, flags)
	return counter, nil
}

// Returns a new performance counter for counting CPU cycles,
// or nil and an error. An error can be returned if the OS is not Linux.
//
// @param user   Specifies whether to count cycles spent in user-space
// @param kernel Specifies whether to count cycles spent in kernel-space
func NewCounter_CpuCycles(user, kernel bool) (*Counter, error) {
	counter, err := newCounterObject()
	if err != nil {
		return nil, err
	}

	var flags uint64
	if !user {
		flags |= FLAG_EXCLUDE_USER
	}
	if !kernel {
		flags |= FLAG_EXCLUDE_KERNEL
	}
	counter.attr.init(TYPE_HARDWARE, HW_CPU_CYCLES, flags)
	return counter, nil
}

// Returns a new performance counter for counting retired instructions,
// or nil and an error. An error can be returned if the OS is not Linux.
//
// @param user   Specifies whether to count instructions executed in user-space
// @param kernel Specifies whether to count instructions executed in kernel-space
func NewCounter_Instructions(user, kernel bool) (*Counter, error) {
	counter, err := newCounterObject()
	if err != nil {
		return nil, err
	}

	var flags uint64
	if !user {
		flags |= FLAG_EXCLUDE_USER
	}
	if !kernel {
		flags |= FLAG_EXCLUDE_KERNEL
	}
	counter.attr.init(TYPE_HARDWARE, HW_INSTRUCTIONS, flags)
	return counter, nil
}

// Returns the thread ID of the calling OS thread
func (c *Counter) Gettid() int {
	if c == nil {
		// It is more comprehensible to panic here,
		// rather than potentially panicing in gettid()
		panic("nil performance-counter object")
	}

	return gettid()
}

// Reads the current value of the performance event counter
func (c *Counter) Read() (n uint64, err error) {
	if c == nil {
		// It is more comprehensible to report an error here,
		// rather than potentially panicing in gettid()
		return 0, errors.New("nil performance-counter object")
	}

	var fd *os.File
	{
		tid := gettid()

		var present bool
		fd, present = c.fd[tid]

		if !present {
			fd, err = c.attr.open(tid)
			if err != nil {
				return
			}

			c.fd[tid] = fd
		}
	}

	var buf [8]byte

	var num_read int
	num_read, err = fd.Read(buf[0:8])
	if err != nil {
		return
	}
	if num_read != 8 {
		panic("expected 8 bytes of data")
	}

	n = *(*uint64)(unsafe.Pointer(&buf[0]))
	return
}

func (c *Counter) Close() error {
	var err error = nil

	for _, file := range c.fd {
		err2 := file.Close()
		if err2 != nil {
			// Report only the 1st error
			if err == nil {
				err = err2
			}
		}
	}

	// Clear 'c.fd'
	c.fd = make(map[int](*os.File))

	return err
}
