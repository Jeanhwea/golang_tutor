package nc0080

import (
	"math"
)

// 买卖股票的最好时机(一)
func maxProfit(prices []int) (ans int) {
	ans = math.MinInt
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
	if ans == math.MinInt {
		ans = 0
	}
	return
}
