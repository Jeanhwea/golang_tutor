package nc0071

// 最长上升子序列 (一)
func LIS(nums []int) (ans int) {
	n := len(nums)
	if n <= 0 {
		return
	}

	dp := [1000]int{} // dp[i] 表示已 nums[i] 结尾的最长升序子序列
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j <= i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
