package lc033

// 搜索旋转排序数组
//
//   1. 要求 O(logN) 算法复杂度
//   2. 采样二分查找法
//   3. 关键是：将区间二分, 左右区间肯定有一个是有序的
//   4. 判断条件 := a[mid] >= a[left] && a[mid] >= a[right] => 左区间有序
//              |= 否则                                    => 右区间有序
//
func search(nums []int, target int) (ans int) {
	ans = -1

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans = mid
			return
		}
		if nums[mid] >= nums[left] && nums[mid] >= nums[right] {
			// 左区间有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右区间有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return
}
