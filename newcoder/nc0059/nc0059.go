package nc0059

const (
	N = 10
)

// N 皇后
func Nqueen(n int) (ans int) {
	count = 0
	forbidden := [N][N]bool{}
	backtrack(forbidden, 0, n)
	ans = count
	return
}

var (
	count int
)

// 回溯法求解
func backtrack(forbidden [N][N]bool, x, n int) {
	if x >= n { // 如果 n 个皇后都放满则输出一种解法
		count++
		return
	}

	for y := 0; y < n; y++ {
		if forbidden[x][y] { // 如果不能落子则继续
			continue
		}

		// 标记不能落子的格子
		var xs, ys []int
		flip := func(x1, y1 int) {
			if x1 >= 0 && x1 < n && y1 >= 0 && y1 < n && !forbidden[x1][y1] {
				xs, ys, forbidden[x1][y1] = append(xs, x1), append(ys, y1), true
			}
		}
		for i := 0; i < n; i++ {
			flip(i, y)
			flip(x, i)
			flip(x+i, y+i)
			flip(x-i, y-i)
			flip(x+i, y-i)
			flip(x-i, y+i)
		}

		backtrack(forbidden, x+1, n)

		// 释放标记的格子
		for i := 0; i < len(xs); i++ {
			forbidden[xs[i]][ys[i]] = false
		}
	}
}
