package lc0094

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
)

func Test_LC0094_01(t *testing.T) {
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
	ans := inorderTraversal(root)
	t.Log(ans)
}
