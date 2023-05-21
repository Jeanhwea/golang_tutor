package nc0073

// 最长回文子串 - (中心扩展法)
func getLongestPalindrome(s string) (ans int) {
	n := len(s)
	for i := 0; i < n; i++ {
		// 奇数回文串情形
		for j := 0; i-j >= 0 && i+j < n; j++ {
			if s[i-j] == s[i+j] {
				length := 2*j + 1
				if length > ans {
					ans = length
				}
			} else {
				break
			}
		}
		// 偶数回文串情形
		if i+1 >= n || s[i] != s[i+1] {
			continue
		}
		for j := 0; i-j >= 0 && i+1+j < n; j++ {
			if s[i-j] == s[i+1+j] {
				length := 2*j + 2
				if length > ans {
					ans = length
				}
			} else {
				break
			}
		}
	}
	return
}

// manacher 算法
