package nc0051

// 半数以上的数字, 选举法
func MoreThanHalfNum_Solution(nums []int) (ans int) {
	count := 0
	for _, v := range nums {
		if count == 0 {
			ans = v
			count++
			continue
		}

		if ans == v {
			count++
			continue
		}

		count--
		if count == 0 {
			ans = v
		}
	}
	return
}
