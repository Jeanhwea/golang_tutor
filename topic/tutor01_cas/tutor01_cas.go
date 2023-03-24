package main

import (
	"fmt"
	"sync/atomic"
)

var (
	readCounter int64
)

func main() {
	atomic.AddInt64(&readCounter, 1)
	atomic.AddInt64(&readCounter, 1)
	fmt.Println(readCounter)
}
