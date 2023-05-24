package nc0083

// 字符串变形
func trans(s string, n int) (ans string) {
	sentence, k := make([]byte, n), n-1
	for i := 0; i < n; i++ {
		if s[i] != ' ' {
			continue
		}

		for j := i - 1; j >= 0 && s[j] != ' '; j-- {
			sentence[k] = flip(s[j])
			k--
		}
		sentence[k] = ' '
		k--
	}

	ans = string(sentence)
	return
}

// 转换单个字符的大小写
func flip(ch byte) (ans byte) {
	if ch >= 'a' && ch <= 'z' {
		ans = byte(int(ch) + ('A' - 'a'))
		return
	}

	if ch >= 'A' && ch <= 'Z' {
		ans = byte(int(ch) + ('a' - 'A'))
		return
	}
	return
}
