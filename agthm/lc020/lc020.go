package lc020

// 有效的括号
//
//   1. 使用一个栈 stack 来存储当前的所有字符
//   2. 如果出现匹配到的括号则出栈
//
func isValid(s string) (ans bool) {
	stack := []byte{}
	match := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range []byte(s) {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) <= 0 {
				ans = false
				return
			}
			top := stack[len(stack)-1]
			if top != match[ch] {
				ans = false
				return
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	if len(stack) == 0 {
		ans = true
		return
	}
	return
}
