package nc0055

// 没有重复项数字的全排列
func permute(nums []int) (ans [][]int) {
	res = [][]int{}
	dfs(nums, 0)
	ans = res
	return
}

var (
	res [][]int
)

func dfs(nums []int, index int) {
	if index == len(nums) {
		res = append(res, append([]int{}, nums...))
		return
	}

	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		dfs(nums, index+1)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
