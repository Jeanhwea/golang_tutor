package nc0054

import (
	"sort"
)

func threeSum(nums []int) (ans [][]int) {
	n := len(nums)
	if n < 3 {
		panic("at least have 3 numbers!")
	}

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		for i > 0 && i < n-2 && nums[i-1] == nums[i] {
			i++
		}

		target := -nums[i]
		for j, k := i+1, n-1; j < k; {
			if nums[j]+nums[k] < target {
				j++
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				val := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, val)
				j++
				for j < k && nums[j-1] == nums[j] {
					j++
				}
			}
		}
	}
	return
}
