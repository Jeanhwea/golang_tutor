package lc0022

// 括号生成
//
//  1. 使用动态规划的解决
//  2. 由于每个序列都是由 (a)b  -- a,b 是子序列, 并且 a, b 可能为空
//
func generateParenthesis(n int) (ans []string) {
	m := map[int][]string{0: {""}}

	for k := 1; k <= n; k++ {
		var vals []string
		for i := 0; i < k; i++ {
			for _, left := range m[i] {
				for _, right := range m[k-i-1] {
					val := "(" + left + ")" + right
					vals = append(vals, val)
				}
			}
		}
		m[k] = vals
	}

	ans = m[n]
	return
}
