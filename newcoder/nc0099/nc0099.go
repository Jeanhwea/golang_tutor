package nc0099

// 顺时针旋转矩阵
func rotateMatrix(mat [][]int, n int) (ans [][]int) {
	for i := 0; i < n; i++ {
		var row []int
		for j := n - 1; j >= 0; j-- {
			row = append(row, mat[j][i])
		}
		ans = append(ans, row)
	}
	return
}

// 原矩阵
// 1 2 3
// 4 5 6
// 7 8 9

// 顺时针旋转后的矩阵
// 7 4 1
// 8 5 2
// 9 6 3
