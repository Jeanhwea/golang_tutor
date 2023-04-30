package lc007

func reverse(x int) (ans int) {
	for x != 0 {
		d := x % 10

		// 最大整数判断条件, 2147483647
		if ans > 214748364 || (ans == 214748364 && d > 7) {
			ans = 0
			return
		}
		// 最小整数判断条件, -2147483648
		if ans < -214748364 || (ans == -214748364 && d < -8) {
			ans = 0
			return
		}

		ans = 10*ans + d
		x = x / 10
	}
	return
}
