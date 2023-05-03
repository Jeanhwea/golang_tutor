package lc003

import (
	"testing"
)

func Test_LC003_01(t *testing.T) {
	str := "abcabcbb"
	ans := lengthOfLongestSubstring(str)
	t.Log(ans)
}

func Test_LC003_02(t *testing.T) {
	str := "pwwkew"
	ans := lengthOfLongestSubstring(str)
	t.Log(ans)
}

func Test_LC003_03(t *testing.T) {
	str := " "
	ans := lengthOfLongestSubstring(str)
	t.Log(ans)
}
