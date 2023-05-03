package lc0415

import (
	"testing"
)

func Test_LC0415_01(t *testing.T) {
	ans := addStrings("190", "23")
	t.Log(ans)
}

func Test_LC0415_02(t *testing.T) {
	ans := addStrings("0", "0")
	t.Log(ans)
}
