package nc0092

// 最长无重复子数组
func maxLength(nums []int) (ans int) {
	// has 记录窗口范围包含的字符
	has, n := make(map[int]bool), len(nums)

	i, j := 0, 0 // i, j 表示窗口的左右指针，窗口的范围为 [i,j)
	for {
		if j < n && has[nums[j]] { // 如果 nums[j] 已经存在
			if ans < j-i {
				ans = j - i
			}

			// 移动左指针来维护 has 数据结构
			for ; nums[i] != nums[j]; i++ {
				has[nums[i]] = false
			}
			i++
		}

		if j >= n {
			if ans < j-i {
				ans = j - i
			}
			break
		}

		has[nums[j]] = true
		j++
	}

	return
}
