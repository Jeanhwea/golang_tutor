package nc0094

// 接雨水
func maxWater(arr []int) (ans int64) {
	n := len(arr)
	left, right := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], arr[i-1])
	}

	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], arr[i+1])
	}

	for i := 0; i < n; i++ {
		ans += int64(max(min(left[i], right[i])-arr[i], 0))
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
