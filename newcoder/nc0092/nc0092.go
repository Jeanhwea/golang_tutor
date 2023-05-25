package nc0092

// 最长无重复子数组
func maxLength(nums []int) (ans int) {
	n := len(nums)
	exists := make(map[int]bool)
	for i, j := 0, 0; j < n; j++ {
		if !exists[nums[j]] {
			exists[nums[j]] = true
			if ans < j-i+1 {
				ans = j - i + 1
			}
			continue
		}

		for ; nums[i] != nums[j]; i++ {
			exists[nums[i]] = false
		}
	}
	return
}
