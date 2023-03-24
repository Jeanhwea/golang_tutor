package tutor02_parallel

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTutor02Parallel01(t *testing.T) {
	main()
}

func TestTutor02Parallel02(t *testing.T) {
	var counter int64
	for i := 0; i < 100; i++ {
		tid := i
		go func() {
			counter = counter + 10
			t.Logf("tid = %v", tid)
		}()
	}
	time.Sleep(3 * time.Second)
	t.Log(counter)
}

func TestTutor02Parallel03(t *testing.T) {
	var counter int64
	for i := 0; i < 100; i++ {
		go func() {
			// CAS
			atomic.AddInt64(&counter, 10)
			// t.Logf("tid = %v", tid)
		}()
	}
	time.Sleep(3 * time.Second)
	t.Log(counter)
}

func TestTutor02Parallel04(t *testing.T) {
	var counter int64
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter = counter + 10
			mu.Unlock()
		}()
	}
	time.Sleep(3 * time.Second)
	t.Log(counter)
}
