package lc020

import (
	"testing"
)

func Test_LC020_01(t *testing.T) {
	s := "()"
	t.Log(isValid(s))
}

func Test_LC020_02(t *testing.T) {
	s := "(]"
	t.Log(isValid(s))
}

func Test_LC020_03(t *testing.T) {
	s := "]"
	t.Log(isValid(s))
}
