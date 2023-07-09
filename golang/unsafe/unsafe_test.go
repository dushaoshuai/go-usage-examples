package unsafe_test

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func ExampleSizeof() {
	// runtime.internal.sys.nih
	type nih struct{}
	// runtime.internal.sys.NotInHeap
	type NotInHeap struct{ _ nih }
	// runtime.lockRankStruct
	type lockRankStruct struct{}
	// runtime.mutex
	type mutex struct {
		lockRankStruct
		key uintptr
	}
	// runtime.timer
	type timer struct {
		pp       uintptr
		when     int64
		period   int64
		f        func(any, uintptr)
		arg      any
		seq      uintptr
		nextwhen int64
		status   atomic.Uint32
	}
	// runtime.pollDesc
	type pollDesc struct {
		_          NotInHeap
		link       *pollDesc
		fd         uintptr
		atomicInfo atomic.Uint32
		rg         atomic.Uintptr
		wg         atomic.Uintptr
		lock       mutex
		closing    bool
		user       uint32
		rseq       uintptr
		rt         timer
		rd         int64
		wseq       uintptr
		wt         timer
		wd         int64
		self       *pollDesc
	}

	// 最开始看这里的代码时没理解，还以为只分配一次，好奇一次能分配多少个，可以保证够用：
	// https://cs.opensource.google/go/go/+/refs/tags/go1.20.5:src/runtime/netpoll.go;l=620;drc=aab8d2b448c4855a4e4a9c2d477671a75828f78b.
	// 发现一次只能分配 17 个 pollDesc。
	// 后面又看懂了，不是只分配一次，是 free list 中没有可用的 pollDesc 时，就会再分配 17 个。
	const pdSize = unsafe.Sizeof(pollDesc{})
	const pollBlockSize = 4 * 1024
	n := pollBlockSize / pdSize
	fmt.Printf("%d / %d = %d\n", pollBlockSize, pdSize, n)

	// Output:
	// 4096 / 240 = 17
}
