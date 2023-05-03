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

func Test_LC020_04(t *testing.T) {
	stack := []byte{}
	stack = append(stack, 'a')
	t.Log(stack)
	stack = append([]byte{'b'}, stack...)
	t.Log(stack)
}
