package lc103

import (
	"testing"
)

func Test_LC103_01(t *testing.T) {
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
