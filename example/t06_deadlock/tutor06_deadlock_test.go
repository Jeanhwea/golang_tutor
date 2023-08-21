package main

import (
	"testing"
	"time"
)

// 非缓存chan 只写不读
func TestTutor06Deadlock01(t *testing.T) {
	ch := make(chan int)
	ch <- 1 // 阻塞, deadlock
}

// 非缓存chan 读在写后
func TestTutor06Deadlock02(t *testing.T) {
	ch := make(chan int)
	ch <- 1 // 阻塞, deadlock
	go func() {
		out := <-ch
		t.Log(out)
	}()
	time.Sleep(time.Second)
}

// 缓存chan, 写超出缓冲区
func TestTutor06Deadlock03(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // 阻塞, deadlock
}

func TestTutor06Deadlock10(t *testing.T) {
	gid := GetGid()
	t.Log(gid)
	for i := 0; i < 10; i++ {
		go func() {
			t.Log(GetGid())
		}()
	}
	time.Sleep(time.Second)
}
