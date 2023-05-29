package nc0097

func solve(n, m int, a []int) (ans []int) {
	m = m % n

	// 全部翻转
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	// 翻转 [0, m-1]
	for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	// 翻转 [m, n-1]
	for i, j := m, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	ans = a
	return
}

func solve2(n, m int, a []int) (ans []int) {
	ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[(i+m)%n] = a[i]
	}
	return
}
