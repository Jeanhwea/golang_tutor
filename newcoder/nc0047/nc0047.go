package nc0047

// 寻找第 K 大的数字
func findKth(nums []int, n, k int) (ans int) {
	if n <= 0 {
		return // no found
	}

	p := nums[0]
	var small, same, large []int
	for _, num := range nums {
		if num < p {
			small = append(small, num)
		} else if num > p {
			large = append(large, num)
		} else {
			same = append(same, num)
		}
	}

	if k <= len(large) {
		ans = findKth(large, len(large), k)
	} else if k <= (len(large) + len(same)) {
		ans = p
	} else {
		ans = findKth(small, len(small), k-len(large)-len(same))
	}
	return
}
