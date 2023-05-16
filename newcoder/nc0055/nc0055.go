package nc0055

// 没有重复项数字的全排列
func permute(nums []int) (ans [][]int) {
	if len(nums) == 1 {
		ans = append(ans, nums)
		return
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		nums[0], nums[i] = nums[i], nums[0]
		num := nums[0]
		for _, perm := range permute(nums[1:]) {
			ans = append(ans, append([]int{num}, perm...))
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return
}
