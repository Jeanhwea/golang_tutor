package nc0079

// 打家劫舍(二)
func rob(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		ans = nums[0]
		return
	}

	dp := make([]int, len(nums))

	// Case 1: 打劫第 0 家
	dp[0], dp[1] = nums[0], nums[0]
	for i := 2; i < n-1; i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	ans = max(ans, dp[n-2])

	// Case 2: 不打劫第 0 家
	dp[0], dp[1] = 0, nums[1]
	for i := 2; i < n; i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	ans = max(ans, dp[n-1])

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
