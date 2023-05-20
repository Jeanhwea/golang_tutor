package nc0070

func minMoney(arr []int, aim int) (ans int) {
	dp := [5001]int{}
	for i := 1; i <= aim; i++ {
		dp[i] = aim + 1
	}
	for i := 1; i <= aim; i++ {
		for _, coin := range arr {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[aim] > aim {
		ans = -1
	} else {
		ans = dp[aim]
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
