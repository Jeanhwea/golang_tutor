package nc0094

import "fmt"

// 接雨水 - 解法一: 动归计算出左右边界后再求值
func maxWater(arr []int) (ans int64) {
	n := len(arr)

	// left[i]  表示 arr[i] 对应的左边界最大值
	// right[i] 表示 arr[i] 对应的右边界最大值
	left, right := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], arr[i-1])
	}

	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], arr[i+1])
	}

	fmt.Printf("left: %v, right: %v\n", left, right)

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

// 双指针计算法
func maxWater2(arr []int) (ans int64) {
	// 维护当前的短板
	leftLower, rightLower := 0, 0

	// 双指针 - 由于短板必须准确，每次移动短板的指针
	i, j := 0, len(arr)-1
	for i <= j {
		if leftLower < rightLower {
			ans += int64(leftLower - min(arr[i], leftLower))
			leftLower = max(leftLower, arr[i])
			i++
		} else {
			ans += int64(rightLower - min(arr[j], rightLower))
			rightLower = max(rightLower, arr[j])
			j--
		}
	}
	return
}
