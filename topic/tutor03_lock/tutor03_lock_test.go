package tutor03_lock

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Benchmark_Mutex(b *testing.B) {
	var lock sync.Mutex
	m := map[int]int{1: 2}

	for i := 0; i < b.N; i++ {
		lock.Lock()
		_ = m[1]
		lock.Unlock()
	}
}

func Benchmark_RWMutex(b *testing.B) {
	var lock sync.RWMutex
	m := map[int]int{1: 2}

	for i := 0; i < b.N; i++ {
		lock.RLock()
		_ = m[1]
		lock.RUnlock()
	}
}

func Benchmark_Atomic(b *testing.B) {
	var ptr atomic.Value
	ptr.Store(map[int]int{1: 2})

	for i := 0; i < b.N; i++ {
		m := ptr.Load().(map[int]int)
		_ = m[1]
	}
}

func Benchmark_Mutex_parallel(b *testing.B) {
	var lock sync.Mutex
	m := map[int]int{1: 2}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			_ = m[1]
			lock.Unlock()
		}
	})
}

func Benchmark_RWMutex_parallel(b *testing.B) {
	var lock sync.RWMutex
	m := map[int]int{1: 2}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.RLock()
			_ = m[1]
			lock.RUnlock()
		}
	})
}

func Benchmark_Atomic_parallel(b *testing.B) {
	var ptr atomic.Value
	ptr.Store(map[int]int{1: 2})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := ptr.Load().(map[int]int)
			_ = m[1]
		}
	})
}
