package nc0093

// 盛水最多的容器
func maxArea(height []int) (ans int) {
	n := len(height)
	if n < 2 {
		return
	}
	for i, j := 0, n-1; i < j; {
		area := (j - i) * min(height[i], height[j])
		if area > ans {
			ans = area
		}
		// 利用木桶原理, 短板的边移动可能会获取到更大的面积
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
