package lc005

import (
	"testing"
)

func Test_LC005_01(t *testing.T) {
	s := "babad"
	ans := longestPalindrome(s)
	t.Log(ans)
}

func Test_LC005_02(t *testing.T) {
	s := "cbbd"
	ans := longestPalindrome(s)
	t.Log(ans)
}

func Test_LC005_11(t *testing.T) {
	s := "abc"
	t.Log(s[0:2])
}
