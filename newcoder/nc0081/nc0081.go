package nc0081

// 买卖股票的最好时机(二)
func maxProfit(prices []int) (ans int) {
	for i := 0; i < len(prices)-1; i++ {
		d := prices[i+1] - prices[i]
		if d > 0 {
			ans += d
		}
	}
	return
}
