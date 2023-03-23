package tutor02_parallel

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	fmt.Println("The program starts ...")

	go func() {
		for {
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("I got scheduled!")
}
