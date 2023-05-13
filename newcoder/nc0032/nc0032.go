package nc0032

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 合并二叉树
func mergeTrees(t1, t2 *TreeNode) (ans *TreeNode) {
	return merge(t1, t2)
}

func merge(root1, root2 *TreeNode) (ans *TreeNode) {
	if root1 == nil && root2 == nil {
		return
	} else if root1 != nil && root2 != nil {
		ans = &TreeNode{
			Val:   root1.Val + root2.Val,
			Left:  merge(root1.Left, root2.Left),
			Right: merge(root1.Right, root2.Right),
		}
		return
	}

	if root1 == nil {
		root1, root2 = root2, root1
	}

	ans = &TreeNode{
		Val:   root1.Val,
		Left:  merge(root1.Left, nil),
		Right: merge(root1.Right, nil),
	}
	return
}
