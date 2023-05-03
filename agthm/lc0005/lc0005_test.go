package lc0005

import (
	"testing"
)

func Test_LC0005_01(t *testing.T) {
	s := "babad"
	ans := longestPalindrome(s)
	t.Log(ans)
}

func Test_LC0005_02(t *testing.T) {
	s := "cbbd"
	ans := longestPalindrome(s)
	t.Log(ans)
}

func Test_LC0005_11(t *testing.T) {
	ss := [10][10]bool{}
	ss[1][1] = true
	t.Log(ss)
}
