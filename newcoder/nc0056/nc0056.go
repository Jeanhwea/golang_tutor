package nc0056

import (
	"sort"
)

func permuteUnique(nums []int) (ans [][]int) {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	vis := make([]int, len(nums))
	ans = [][]int{}
	temp := []int{}
	perm(&ans, nums, temp, vis)
	return ans
}

func perm(ans *[][]int, nums []int, choose []int, visited []int) {
	if len(choose) == len(nums) {
		vals := make([]int, len(choose))
		copy(vals, choose)
		*ans = append(*ans, vals)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] > 0 {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && visited[i-1] == 0 {
			continue
		}
		visited[i] = 1
		choose = append(choose, nums[i])
		perm(ans, nums, choose, visited)
		visited[i] = 0
		choose = choose[:len(choose)-1]
	}
}
