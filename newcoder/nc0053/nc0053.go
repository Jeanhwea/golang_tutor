package nc0053

// 缺失的第一个正数
//
//  该题有点技巧性: n = len(nums)
//  假设我们有一个哈希表 m
//  情况一: ans >= 0 && ans <= n, 则遍历 1..n 出现 m[i] 不存在就是最小正数
//  情况二: ans >= 0 && ans <= n, 则遍历 1..n 出现 m[.] 都存在, 则最小正数为 n+1
//
//  该题的精髓是将 nums 当成一个 hashmap 使用来节约空间,
//    1. 其中 abs(val) 表示对应的数
//    2. val 的符号表示是否存在, 负号表示存在
//
func minNumberDisappeared(nums []int) (ans int) {
	n := len(nums)

	// 先处理负数, 将负数都置成 n+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 将 nums 数组当成 hashmap 使用, 如果 nums[x-1] < 0, 表示 x 存在
	for i := 0; i < n; i++ {
		x := abs(nums[i])
		if x <= n && nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		}
	}

	// 遍历 hashmap, 找到第一个缺失的正数
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			ans = i + 1
			return
		}
	}
	ans = n + 1

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
