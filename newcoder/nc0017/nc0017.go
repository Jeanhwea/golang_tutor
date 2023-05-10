package nc0017

// 二分查找-I
func search(nums []int, target int) (ans int) {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans = mid
			return
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	ans = -1
	return
}
