package lc0053

import "math"

// 最大子数组和
func maxSubArray(nums []int) (ans int) {
	ans = math.MinInt
	sum := 0
	for _, v := range nums {
		sum += v
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return
}
