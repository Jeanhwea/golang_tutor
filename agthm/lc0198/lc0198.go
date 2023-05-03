package lc0198

// 打家劫舍
//
//    动态规划: dp[i] 表示打劫到 nums[i] 家偷窃到的最高金额
//    dp[i] := max(nums[i] + dp[i-2], dp[i-1])
//          |= nums[0] iff i == 0
//          |= max(nums[0],nums[1]) iff i == 1
//
func rob(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	const N = 100
	dp := [N]int{}

	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	ans = dp[n-1]
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
