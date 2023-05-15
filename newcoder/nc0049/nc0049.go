package nc0049

func solve(s string) (ans int) {
	var op []byte
	var nums []int
	var num, sign int
	for i, n := 0, len(s); i < n; i++ {
		switch ch := s[i]; ch {
		case '(', '+':
			op = append(op, ch)
		default:
			for num := 0; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
				num = 10*num + int(s[i]-'0')
			}
			nums = append(nums, sign*num)
		}
	}
	return
}
