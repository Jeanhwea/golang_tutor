package nc0083

// 字符串变形
func trans(s string, n int) (ans string) {
	sentence, i, j, k := make([]byte, n), 0, -1, n-1
	for {
		for i < n && s[i] != ' ' {
			i++
		}

		for p := i - 1; p > j; p-- {
			sentence[k] = flip(s[p])
			k--
		}
		j = i
		if k >= 0 {
			sentence[k] = ' '
			k--
		}
		i++
		if i >= n {
			break
		}
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
