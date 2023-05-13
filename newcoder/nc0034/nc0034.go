package nc0034

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 判断是否是搜索二叉树
func isValidBST(root *TreeNode) bool {
	prev, isBST = nil, true
	visit(root)
	return isBST
}

var (
	prev  *TreeNode
	isBST bool
)

func visit(root *TreeNode) {
	if root == nil {
		return
	}
	visit(root.Left)
	if prev != nil && prev.Val > root.Val {
		isBST = false
		return
	} else {
		prev = root
	}
	visit(root.Right)
}
