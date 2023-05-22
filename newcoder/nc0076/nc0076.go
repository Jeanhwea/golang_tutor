package nc0076

// 正则表达式匹配

// 递归解法
func match(str string, pat string) (ans bool) {
	n, m := len(str), len(pat)

	switch m {
	case 0:
		return n == 0
	case 1:
		if n != 1 {
			return false
		}
		return pat[0] == '.' || pat[0] == str[0]
	default: // m >= 2
		if pat[1] == '*' {
			if n == 0 {
				return match(str, pat[2:])
			}

			if pat[0] == '.' || pat[0] == str[0] {
				return match(str[1:], pat) || match(str, pat[2:])
			} else {
				return match(str, pat[2:])
			}
		} else {
			if n == 0 {
				return false
			}
			if pat[0] == '.' || pat[0] == str[0] {
				return match(str[1:], pat[1:])
			}
		}
	}

	return false
}
