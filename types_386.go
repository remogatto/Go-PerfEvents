// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types.go

package perf

const _SYS_PERF_OPEN = 0x150

// Event types
const (
	TYPE_HARDWARE   = 0x0
	TYPE_SOFTWARE   = 0x1
	TYPE_TRACEPOINT = 0x2
	TYPE_HW_CACHE   = 0x3
	TYPE_RAW        = 0x4
	TYPE_BREAKPOINT = 0x5
)

/*
 * Generalized hardware events
 */
const (
	HW_CPU_CYCLES          = 0x0
	HW_INSTRUCTIONS        = 0x1
	HW_CACHE_REFERENCES    = 0x2
	HW_CACHE_MISSES        = 0x3
	HW_BRANCH_INSTRUCTIONS = 0x4
	HW_BRANCH_MISSES       = 0x5
	HW_BUS_CYCLES          = 0x6
)

/*
 * Generalized hardware cache events:
 *
 *       { L1-D, L1-I, LLC, ITLB, DTLB, BPU, NODE } x
 *       { read, write, prefetch } x
 *       { accesses, misses }
 */
const (
	HW_CACHE_L1D           = 0x0
	HW_CACHE_L1I           = 0x1
	HW_CACHE_LL            = 0x2
	HW_CACHE_DTLB          = 0x3
	HW_CACHE_ITLB          = 0x4
	HW_CACHE_BPU           = 0x5
	HW_CACHE_NODE          = 0x6
	HW_CACHE_OP_READ       = 0x0
	HW_CACHE_OP_WRITE      = 0x1
	HW_CACHE_OP_PREFETCH   = 0x2
	HW_CACHE_RESULT_ACCESS = 0x0
	HW_CACHE_RESULT_MISS   = 0x1
)

/*
 * Special "software" events provided by the kernel, even if the hardware
 * does not support performance events. These events measure various
 * physical and sw events of the kernel (and allow the profiling of them as
 * well):
 */
const (
	SW_CPU_CLOCK        = 0x0
	SW_TASK_CLOCK       = 0x1
	SW_PAGE_FAULTS      = 0x2
	SW_CONTEXT_SWITCHES = 0x3
	SW_CPU_MIGRATIONS   = 0x4
	SW_PAGE_FAULTS_MIN  = 0x5
	SW_PAGE_FAULTS_MAJ  = 0x6
	SW_ALIGNMENT_FAULTS = 0x7
	SW_EMULATION_FAULTS = 0x8
)

const (
	FLAG_DISABLED = 1 << iota
	FLAG_INHERIT
	FLAG_PINNED
	FLAG_EXCLUSIVE
	FLAG_EXCLUDE_USER
	FLAG_EXCLUDE_KERNEL
	FLAG_EXCLUDE_HV
	FLAG_EXCLUDE_IDLE
	FLAG_MMAP
	FLAG_COMM
	FLAG_FREQ
	FLAG_INHERIT_STAT
	FLAG_ENABLE_ON_EXEC
	FLAG_TASK
	FLAG_WATERMARK
)

// Size of struct Attr
const _ATTR_SIZE = 72

type Attr struct {
	Type                     uint32
	Size                     uint32
	Config                   uint64
	Sample_periodOrFreq      uint64
	Sample_type              uint64
	Read_format              uint64
	Flags                    uint64
	Wakeup_eventsOrWatermark uint32
	Bp_type                  uint32
	Bp_addr                  uint64
	Bp_len                   uint64
}
