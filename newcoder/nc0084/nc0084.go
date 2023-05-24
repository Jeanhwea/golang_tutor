package nc0084

// 最长公共前缀
func longestCommonPrefix(strs []string) (ans string) {
	n := len(strs)
	if n == 0 {
		return
	}
	var prefix []byte
	var letter byte
	for j := 0; ; j++ {
		for k := 0; k < n; k++ {
			if j >= len(strs[k]) {
				goto done
			} else {
				if k == 0 {
					letter = strs[k][j]
				} else {
					if strs[k][j] != letter {
						goto done
					}
				}
			}
		}
		prefix = append(prefix, letter)
	}
done:

	ans = string(prefix)
	return
}
