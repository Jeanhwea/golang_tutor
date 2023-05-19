package nc0063

func jumpFloor(n int) (ans int) {
	dp := [40]int{1, 2}
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	ans = dp[n-1]
	return
}
