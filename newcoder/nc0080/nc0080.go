package nc0080

// 买卖股票的最好时机(一)
func maxProfit(prices []int) (ans int) {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		sum += prices[i+1] - prices[i]
		if sum < 0 {
			sum = 0
		}
		if sum > ans {
			ans = sum
		}
	}
	return
}
