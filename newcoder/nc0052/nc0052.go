package nc0052

// 数组中只出现一次的两个数字
func FindNumsAppearOnce(nums []int) (ans []int) {
	out := 0
	for _, v := range nums {
		out ^= v
	}

	mask := 1
	for (out & mask) == 0 {
		mask = mask << 1
	}

	a, b := 0, 0
	for _, v := range nums {
		if (v & mask) == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}

	if a > b {
		a, b = b, a
	}

	ans = []int{a, b}
	return
}
