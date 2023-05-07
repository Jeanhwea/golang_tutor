// https://leetcode.cn/problems/two-sum/

package lc0001

func twoSum(nums []int, target int) (ans []int) {
	cache := make(map[int]int)
	for i, a := range nums {
		if j, ok := cache[a]; ok {
			return []int{j, i}
		}
		cache[target-a] = i
	}
	return
}
