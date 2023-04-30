package lc015

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

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		target := -nums[i]
		for p, q := i+1, n-1; p < q; {
			if nums[p]+nums[q] > target {
				q--
			} else if nums[p]+nums[q] < target {
				p++
			} else {
				ans = append(ans, []int{nums[i], nums[p], nums[q]})
				p++
				for p < q && nums[p-1] == nums[p] {
					p++
				}
			}
		}
	}
	return
}
