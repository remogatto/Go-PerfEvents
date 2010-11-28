package perf

import (
	"os"
	"unsafe"
)


type PerfCounter struct {
	attr Attr
	fd   map[int](*os.File) // File descriptors for each OS thread, initially empty
}


func (attr *Attr) init_HW(event uint, exclude_user bool, exclude_kernel bool) {
	attr.Type = TYPE_HARDWARE
	attr.Size = ATTR_SIZE
	attr.Config = uint64(event)

	var flags uint64 = 0
	if exclude_user {
		flags |= FLAG_EXCLUDE_USER
	}
	if exclude_kernel {
		flags |= FLAG_EXCLUDE_KERNEL
	}
	flags |= FLAG_EXCLUDE_HV
	flags |= FLAG_EXCLUDE_IDLE
	attr.Flags = flags
}

func (attr *Attr) open(pid int) (counter *os.File, err os.Error) {
	return sys_perf_counter_open(attr, /*pid*/ pid, /*cpu*/ -1, /*group_fd*/ -1, /*flags*/ 0)
}


// Returns a new performance counter for counting CPU cycles,
// or nil and an error in case of a failure. A failure can occur
// if the OS is not Linux.
//
// @param user   Specifies whether to count cycles spent in user-space
// @param kernel Specifies whether to count cycles spent in kernel-space
func NewCounter_CpuCycles(user, kernel bool) (*PerfCounter, os.Error) {
	counter, err := newPerfCounterObject()
	if err != nil {
		return nil, err
	}

	counter.attr.init_HW(HW_CPU_CYCLES, !user, !kernel)
	return counter, nil
}

// Returns a new performance counter for counting retired instructions,
// or nil and an error in case of a failure. A failure can occur
// if the OS is not Linux.
//
// @param user   Specifies whether to count instructions executed in user-space
// @param kernel Specifies whether to count instructions executed in kernel-space
func NewCounter_Instructions(user, kernel bool) (*PerfCounter, os.Error) {
	counter, err := newPerfCounterObject()
	if err != nil {
		return nil, err
	}

	counter.attr.init_HW(HW_INSTRUCTIONS, !user, !kernel)
	return counter, nil
}

// Returns the thread ID of the calling OS thread
func (c *PerfCounter) Gettid() int {
	if c == nil {
		// It is more comprehensible to panic here,
		// rather than potentially panicing in gettid()
		panic("nil performance-counter object")
	}

	return gettid()
}

// Reads the current value of the performance counter
func (c *PerfCounter) Read() (n uint64, err os.Error) {
	if c == nil {
		// It is more comprehensible to report an error here,
		// rather than potentially panicing in gettid()
		return 0, os.NewError("nil performance-counter object")
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

func (c *PerfCounter) Close() os.Error {
	var err os.Error = nil

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
