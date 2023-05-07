package main

import "fmt"

// go build -o a.out
// ./a.out < in.txt
func main() {
	var n int
	nums := [10]int{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(nums)
}
