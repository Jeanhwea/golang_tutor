package nc0067

// 不同路径的数目 (一)
func uniquePaths(n, m int) (ans int) {
	if n > m {
		n, m = m, n
	}
	ans = 1
	for i := 1; i < n; i++ {
		ans = ans * (m + i - 1) / i
	}
	return
}
