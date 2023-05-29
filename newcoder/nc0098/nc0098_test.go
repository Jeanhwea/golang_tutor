package nc0098

import (
	"testing"
)

func Test_NC0098_01(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	ans := spiralOrder(matrix)
	t.Log(ans)
}

func Test_NC0098_02(t *testing.T) {
	matrix := [][]int{}
	ans := spiralOrder(matrix)
	t.Log(ans)
}
