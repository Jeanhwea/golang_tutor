package tutor04_channel

import (
	"testing"
	"time"
)

func TestTutor04Channel01(t *testing.T) {
	ch := make(chan int)
	go func() {
		for n := range ch {
			t.Log(n)
		}
	}()

	for i := 0; i < 3; i++ {
		ch <- i
	}
	time.Sleep(3 * time.Second)
}

func TestTutor04Channel02(t *testing.T) {
	ch := make(chan int)
	go func() {
		for n := range ch {
			t.Logf("get(%v)\n", n)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
		t.Logf("send(%v)\n", i)
	}
	time.Sleep(5 * time.Second)
}

func TestTutor04Channel03(t *testing.T) {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go func() {
		for {
			<-chC
			t.Log("A")
			chA <- 1
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			<-chA
			t.Log("B")
			chB <- 1
		}
	}()
	go func() {
		for {
			<-chB
			t.Log("C")
			chC <- 1
		}
	}()

	chC <- 1
	time.Sleep(3 * time.Second)
}
