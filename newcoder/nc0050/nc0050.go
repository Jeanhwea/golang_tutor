package nc0050

func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := m[nums[i]]; ok {
			ans = []int{j + 1, i + 1}
			return
		} else {
			m[target-nums[i]] = i
		}
	}
	return
}
