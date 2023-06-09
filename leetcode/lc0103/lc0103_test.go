package lc0103

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
)

func Test_LC0103_01(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	ans := zigzagLevelOrder(root)
	t.Log(ans)
}
