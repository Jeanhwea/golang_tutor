package nc0026

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0026_01(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}
	ans := levelOrder(root)
	assert.Equal(t, [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}, ans)
}

func TestNc002601(t *testing.T) {
	var root *TreeNode
	assert.Nil(t, levelOrder(root))
}
