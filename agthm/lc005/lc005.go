package lc005

func longestPalindrome(s string) (ans string) {
	n := len(s)
	for k := n; k > 0; k-- {
		for i := 0; i <= n-k; i++ {
			found := true
			for p, q := i, i+k-1; p < q; p, q = p+1, q-1 {
				if s[p] != s[q] {
					found = false
					break
				}
			}
			if found {
				ans = s[i : i+k]
				return
			}
		}
	}
	return
}
