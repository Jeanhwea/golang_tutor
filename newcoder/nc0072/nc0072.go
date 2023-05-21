package nc0072

import "math"

// 连续子数组的最大和
func FindGreatestSumOfSubArray(nums []int) (ans int) {
	n := len(nums)
	ans = math.MinInt
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return
}
