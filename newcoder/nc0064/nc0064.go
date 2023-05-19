package nc0064

// 最小花费爬楼梯
func minCostClimbingStairs(cost []int) (ans int) {
	n := len(cost)
	if n < 2 {
		ans = 0
		return
	}

	dp := make([]int, n+1, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	ans = dp[n]
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
