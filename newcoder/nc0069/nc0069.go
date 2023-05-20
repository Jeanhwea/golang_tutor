package nc0069

// 把数字翻译成字符串
func solve(nums string) (ans int) {
	n := len(nums)
	dp := [90]int{} // 表示 nums[:i+1] 可以翻译成字符串的计数
	for i := 0; i < n; i++ {
		if i == 0 {
			if isValid(nums[i : i+1]) {
				dp[i] = 1
			}
		} else if i == 1 {
			if isValid(nums[i : i+1]) {
				dp[i] += dp[i-1]
			}
			if isValid(nums[i-1 : i+1]) {
				dp[i] += 1
			}
		} else {
			if isValid(nums[i : i+1]) {
				dp[i] += dp[i-1]
			}
			if isValid(nums[i-1 : i+1]) {
				dp[i] += dp[i-2]
			}
		}
	}
	ans = dp[n-1]
	return
}

func isValid(str string) bool {
	if len(str) == 1 && str[0] != '0' {
		return true
	}

	if len(str) == 2 {
		if str[0] == '1' {
			return true
		} else if str[0] == '2' && str[1] >= '0' && str[1] <= '6' {
			return true
		}
	}

	return false
}
