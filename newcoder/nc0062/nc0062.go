package nc0062

func Fibonacci(n int) (ans int) {
	fib := make([]int, n+1, n+1)
	fib[0], fib[1] = 1, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	ans = fib[n-1]
	return
}
