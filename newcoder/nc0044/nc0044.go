package nc0044

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	var top byte
	for _, ch := range []byte(s) {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) <= 0 {
				return false
			}
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if m[ch] != top {
				return false
			}
		}
	}

	return len(stack) == 0
}
