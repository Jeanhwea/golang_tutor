package lc0322

import (
	"math"
)

// 零钱兑换
//
//    广度优先搜索
//    dp[n] := 1 + min(dp[n-c]), c belongs to coins
//          |= 0  iff n == 0
//          |= -1 iff n < 0
//
func coinChange(coins []int, amount int) (ans int) {
	memo = make(map[int]int)
	ans = bfs(coins, amount)
	return
}

var memo map[int]int

func bfs(coins []int, amount int) (ans int) {
	if val, ok := memo[amount]; ok {
		ans = val
		return
	}

	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	ans = math.MaxInt
	for _, coin := range coins {
		cnt := bfs(coins, amount-coin)
		if cnt >= 0 {
			ans = min(ans, cnt+1)
		}
	}

	if ans == math.MaxInt {
		ans = -1
	}
	memo[amount] = ans
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
