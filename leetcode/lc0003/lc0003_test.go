package lc0003

import (
	"testing"
)

func Test_LC0003_01(t *testing.T) {
	str := "abcabcbb"
	ans := lengthOfLongestSubstring(str)
	t.Log(ans)
}

func Test_LC0003_02(t *testing.T) {
	str := "pwwkew"
	ans := lengthOfLongestSubstring(str)
	t.Log(ans)
}

func Test_LC0003_03(t *testing.T) {
	str := " "
	ans := lengthOfLongestSubstring(str)
	t.Log(ans)
}
