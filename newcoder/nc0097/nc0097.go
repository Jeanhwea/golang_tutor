package nc0097

func solve(n, m int, a []int) (ans []int) {
	ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[(i+m)%n] = a[i]
	}
	return
}
