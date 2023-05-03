package lc0015

import (
	"sort"
)

// 三数之和
//
//   1. 固定一个下标 i, 将三数之和转化成两数之和问题
//   2. 注意题目中需要去重
//
func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)

	for start := 0; start < n-2; start++ {
		if start > 0 && nums[start-1] == nums[start] {
			continue
		}

		target := -nums[start]
		for i, j := start+1, n-1; i < j; {
			if nums[i]+nums[j] > target {
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				ans = append(ans, []int{nums[start], nums[i], nums[j]})
				i++
				for i < j && nums[i-1] == nums[i] {
					i++
				}
			}
		}
	}
	return
}
