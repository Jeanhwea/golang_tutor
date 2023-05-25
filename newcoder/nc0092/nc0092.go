package nc0092

// 最长无重复子数组
func maxLength(nums []int) (ans int) {
	ok, n := make(map[int]bool), len(nums)

	i, j := 0, 0
	for {
		if j < n && ok[nums[j]] {
			if ans < j-i {
				ans = j - i
			}
			for ; nums[i] != nums[j]; i++ {
				ok[nums[i]] = false
			}
			i++
		}

		if j >= n {
			if ans < j-i {
				ans = j - i
			}
			break
		}

		ok[nums[j]] = true
		j++
	}

	return
}
