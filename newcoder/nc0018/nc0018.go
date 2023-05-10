package nc0018

// 二维数组中的查找
func Find(target int, array [][]int) (ans bool) {
	n, m := len(array), len(array[0])
	for i, j := n-1, 0; i >= 0 && i < n && j >= 0 && j < m; {
		if array[i][j] == target {
			ans = true
			return
		} else if array[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return
}
