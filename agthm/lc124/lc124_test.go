package lc124

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LC124_01(t *testing.T) {
	root := &TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	ans := maxPathSum(root)
	assert.Equal(t, 42, ans)
}

func Test_LC124_02(t *testing.T) {
	root := &TreeNode{
		Val:  9,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{
			Val:  -3,
			Left: &TreeNode{Val: -6},
			Right: &TreeNode{Val: 2,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  -6,
						Left: &TreeNode{Val: -6},
					},
					Right: &TreeNode{
						Val: -6,
					},
				},
			},
		},
	}
	ans := maxPathSum(root)
	assert.Equal(t, 16, ans)
}
