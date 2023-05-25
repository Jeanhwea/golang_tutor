package nc0091

// 翻转字符串
func solve(str string) string {
	res := []byte(str)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
