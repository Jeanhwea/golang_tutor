package lc0020

import (
	"testing"
)

func Test_LC0020_01(t *testing.T) {
	s := "()"
	t.Log(isValid(s))
}

func Test_LC0020_02(t *testing.T) {
	s := "(]"
	t.Log(isValid(s))
}

func Test_LC0020_03(t *testing.T) {
	s := "]"
	t.Log(isValid(s))
}

func Test_LC0020_04(t *testing.T) {
	stack := []byte{}
	stack = append(stack, 'a')
	t.Log(stack)
	stack = append([]byte{'b'}, stack...)
	t.Log(stack)
}
