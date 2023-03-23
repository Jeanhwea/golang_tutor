package main

import (
	"sync/atomic"
)

var (
	readCounter int64 = 1
)

func main() {
	atomic.AddInt64(&readCounter, 1)
}
