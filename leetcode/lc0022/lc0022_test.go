package lc0022

import (
	"testing"
)

// ((())) (()()) (())() ()(()) ()()()
func Test_LC0022_01(t *testing.T) {
	t.Log(generateParenthesis(3))
}

// [(((()))),((()())),((())()),((()))(),(()(())),(()()()),(()())(),(())(()),(())()(),()((())),()(()()),()(())(),()()(()),()()()()]
// [(((()))) ((()())) ((())()) ((()))() (()(())) (()()()) (()())() (())()() ()((())) ()(()()) ()(())() ()()(()) ()()()()]
func Test_LC0022_02(t *testing.T) {
	t.Log(generateParenthesis(4))
}

func Test_LC0022_03(t *testing.T) {
	t.Log(generateParenthesis(2))
}

func Test_LC0022_04(t *testing.T) {
	t.Log(generateParenthesis(1))
}
