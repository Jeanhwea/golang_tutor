package nc0099

import (
	"testing"
)

func Test_NC0099_01(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	ans := rotateMatrix(matrix, 3)
	t.Log(ans)
}
