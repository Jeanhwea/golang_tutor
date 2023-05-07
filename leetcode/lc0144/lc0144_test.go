package lc0144

import (
	"testing"
)

func Test_LC0144_01(t *testing.T) {
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
	ans := preorderTraversal(root)
	t.Log(ans)
}
