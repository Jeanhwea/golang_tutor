package nc0078

func rob(nums []int) (ans int) {
	// dp[i] 表示打劫 nums[0:i+1] 范围房间的最大金额
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = max(nums[0], nums[1])
		} else {
			dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		}
	}

	ans = dp[len(nums)-1]
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
