package nc0019

// 寻找峰值: 也可以使用二分法: 如 nums[mid] < nums[mid+1]
func findPeakElement(nums []int) (ans int) {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	ans = left
	return
}

// 暴力解法
func findPeakElement2(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if i-1 >= 0 && i+1 < n && nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			ans = i
			return
		}

		if i == 0 && i+1 < n && nums[i] > nums[i+1] {
			ans = i
			return
		}

		if i == n-1 && i-1 >= 0 && nums[i] > nums[i-1] {
			ans = i
			return
		}

		if i-1 < 0 && i+1 >= n {
			ans = i
			return
		}
	}
	return
}
