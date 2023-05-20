package nc0067

// 不同路径的数目 (一)
func uniquePaths(n, m int) (ans int) {
	ans = 1
	for i := 1; i < n; i++ {
		ans = ans * (m + i - 1) / i
	}
	return
}
