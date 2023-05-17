package nc0056

import (
	"sort"
)

func permuteUnique(nums []int) (ans [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	ans = [][]int{}
	temp := []int{}
	perm(&ans, nums, temp, visited)
	return ans
}

func perm(ans *[][]int, nums []int, choose []int, visited []bool) {
	if len(choose) == len(nums) {
		vals := make([]int, len(choose))
		copy(vals, choose)
		*ans = append(*ans, vals)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && !visited[i-1] {
			continue
		}
		visited[i], choose = true, append(choose, nums[i])
		perm(ans, nums, choose, visited)
		visited[i], choose = false, choose[:len(choose)-1]
	}
}
