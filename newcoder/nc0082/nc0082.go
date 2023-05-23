package nc0082

// 买卖股票的最好时机(三)
func maxProfit(prices []int) (ans int) {
	// dp[i][5] 表示第 i 天处于状态时的最大利润
	//
	//  数字状态含义如下:
	//       0 - 未买入
	//       1 - 第一次买入
	//       2 - 第一次卖出
	//       3 - 第二次买入
	//       4 - 第二次卖出
	//

	n := len(prices)
	dp := make([][5]int, n)
	for i := 0; i < n; i++ {
		dp[i][0] = -99999
		dp[i][1] = -99999
		dp[i][2] = -99999
		dp[i][3] = -99999
		dp[i][4] = -99999
	}
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	ans = max(dp[n-1][2], dp[n-1][4])
	ans = max(0, ans)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
