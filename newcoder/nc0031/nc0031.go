package nc0031

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 对称的二叉树
func isSymmetrical(root *TreeNode) (ans bool) {
	if root == nil {
		return true
	}

	return symTree(root.Left, root.Right)
}

// 判断 root1, root2 是否为对称树
func symTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 != nil && root1.Val == root2.Val {
		return symTree(root1.Left, root2.Right) && symTree(root1.Right, root2.Left)
	}

	return false
}
