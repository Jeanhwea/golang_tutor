package nc0049

func solve(s string) (ans int) {
	ans, _ = eval(s, 0)
	return
}

func eval(s string, index int) (val, i int) {
	var opts []byte // 操作符栈
	var nums []int  // 操作数

	pushNumber := func(num int) {
		if len(opts) > 0 && opts[len(opts)-1] == '*' {
			lhs := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			opts = opts[:len(opts)-1]
			num = lhs * num
		}
		nums = append(nums, num)
	}

	var num int
	i = index
outter:
	for n := len(s); i < n; i++ {
		switch ch := s[i]; ch {
		case '(': // 如果遇到左括号需要递归处理
			num, i = eval(s, i+1)
			if s[i] != ')' {
				panic("missing right parenthesis!")
			}
			pushNumber(num)
		case ')': // 右括号跳出解析
			break outter
		case '+', '-', '*':
			opts = append(opts, ch)
		case ' ':
			continue
		default: // 解析数字
			num = 0
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = 10*num + int(s[i]-'0')
				i++
			}
			pushNumber(num)
			i--
		}
	}

	val = nums[0]
	for k := 0; k < len(opts); k++ {
		if opts[k] == '+' {
			val += nums[k+1]
		} else if opts[k] == '-' {
			val += -nums[k+1]
		}
	}

	return
}
