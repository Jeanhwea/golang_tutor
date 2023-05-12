package nc0021

// 寻找旋转数组的最小值
func minNumberInRotateArray(rotateArray []int) (ans int) {
	left, right := 0, len(rotateArray)-1
	for left < right {
		mid := left + (right-left)/2
		if rotateArray[mid] > rotateArray[right] {
			left = mid + 1
		} else if rotateArray[mid] < rotateArray[right] {
			right = mid
		} else {
			right--
		}
	}
	ans = rotateArray[left]
	return
}
