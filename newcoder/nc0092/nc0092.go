package nc0092

// 最长无重复子数组
func maxLength(nums []int) (ans int) {
	// has 记录窗口范围包含的字符
	has, n := make(map[int]bool), len(nums)

	i, j := 0, 0 // i, j 表示窗口的左右指针，窗口的范围为 [i,j)
	for {
		// 窗口扩张: 右移动右指针
		for j < n && !has[nums[j]] {
			has[nums[j]] = true
			j++
		}

		// 维护最长的结果值
		if ans < j-i {
			ans = j - i
		}

		// 数组遍历完成后跳出循环
		if j >= n {
			break
		}

		// 窗口收缩: 右移动左指针
		for i < j && nums[i] != nums[j] {
			has[nums[i]] = false
			i++
		}

		i++
		j++
	}

	return
}
