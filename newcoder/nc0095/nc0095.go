package nc0095

// 分糖果
func candy(scores []int) (ans int) {
	n := len(scores)
	nums := make([]int, n)

	// 初始化每个人发一个糖果
	for i := 0; i < n; i++ {
		nums[i] = 1
	}

	// 从左到右进行修正
	for i := 1; i < n; i++ {
		if scores[i-1] < scores[i] {
			nums[i] = nums[i-1] + 1
		}
	}

	// 从右到左进行修正
	for i := n - 2; i >= 0; i-- {
		if scores[i] > scores[i+1] && nums[i] <= nums[i+1] {
			nums[i] = nums[i+1] + 1
		}
	}

	// 累计求和得到结果值
	for i := 0; i < n; i++ {
		ans += nums[i]
	}

	return
}
