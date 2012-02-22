package perf

//#include <linux/unistd.h>
//#include <linux/perf_event.h>
import "C"

const _SYS_PERF_OPEN = C.__NR_perf_event_open

// Event types
const (
	TYPE_HARDWARE          = C.PERF_TYPE_HARDWARE
	TYPE_SOFTWARE          = C.PERF_TYPE_SOFTWARE
	TYPE_TRACEPOINT        = C.PERF_TYPE_TRACEPOINT
	TYPE_HW_CACHE          = C.PERF_TYPE_HW_CACHE
	TYPE_RAW               = C.PERF_TYPE_RAW
	TYPE_BREAKPOINT        = C.PERF_TYPE_BREAKPOINT
)

/*
 * Generalized hardware events
 */
const (
	HW_CPU_CYCLES          = C.PERF_COUNT_HW_CPU_CYCLES
	HW_INSTRUCTIONS        = C.PERF_COUNT_HW_INSTRUCTIONS
	HW_CACHE_REFERENCES    = C.PERF_COUNT_HW_CACHE_REFERENCES
	HW_CACHE_MISSES        = C.PERF_COUNT_HW_CACHE_MISSES
	HW_BRANCH_INSTRUCTIONS = C.PERF_COUNT_HW_BRANCH_INSTRUCTIONS
	HW_BRANCH_MISSES       = C.PERF_COUNT_HW_BRANCH_MISSES
	HW_BUS_CYCLES          = C.PERF_COUNT_HW_BUS_CYCLES
)

/*
 * Generalized hardware cache events:
 *
 *       { L1-D, L1-I, LLC, ITLB, DTLB, BPU, NODE } x
 *       { read, write, prefetch } x
 *       { accesses, misses }
 */
const (
	HW_CACHE_L1D           = C.PERF_COUNT_HW_CACHE_L1D
	HW_CACHE_L1I           = C.PERF_COUNT_HW_CACHE_L1I
	HW_CACHE_LL            = C.PERF_COUNT_HW_CACHE_LL
	HW_CACHE_DTLB          = C.PERF_COUNT_HW_CACHE_DTLB
	HW_CACHE_ITLB          = C.PERF_COUNT_HW_CACHE_ITLB
	HW_CACHE_BPU           = C.PERF_COUNT_HW_CACHE_BPU
	HW_CACHE_NODE          = C.PERF_COUNT_HW_CACHE_NODE
	HW_CACHE_OP_READ       = C.PERF_COUNT_HW_CACHE_OP_READ
	HW_CACHE_OP_WRITE      = C.PERF_COUNT_HW_CACHE_OP_WRITE
	HW_CACHE_OP_PREFETCH   = C.PERF_COUNT_HW_CACHE_OP_PREFETCH
	HW_CACHE_RESULT_ACCESS = C.PERF_COUNT_HW_CACHE_RESULT_ACCESS
	HW_CACHE_RESULT_MISS   = C.PERF_COUNT_HW_CACHE_RESULT_MISS
)

/*
 * Special "software" events provided by the kernel, even if the hardware
 * does not support performance events. These events measure various
 * physical and sw events of the kernel (and allow the profiling of them as
 * well):
 */
const (
	SW_CPU_CLOCK           = C.PERF_COUNT_SW_CPU_CLOCK
	SW_TASK_CLOCK          = C.PERF_COUNT_SW_TASK_CLOCK
	SW_PAGE_FAULTS         = C.PERF_COUNT_SW_PAGE_FAULTS
	SW_CONTEXT_SWITCHES    = C.PERF_COUNT_SW_CONTEXT_SWITCHES
	SW_CPU_MIGRATIONS      = C.PERF_COUNT_SW_CPU_MIGRATIONS
	SW_PAGE_FAULTS_MIN     = C.PERF_COUNT_SW_PAGE_FAULTS_MIN
	SW_PAGE_FAULTS_MAJ     = C.PERF_COUNT_SW_PAGE_FAULTS_MAJ
	SW_ALIGNMENT_FAULTS    = C.PERF_COUNT_SW_ALIGNMENT_FAULTS
	SW_EMULATION_FAULTS    = C.PERF_COUNT_SW_EMULATION_FAULTS
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
