package nc0035

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func isCompleteTree(root *TreeNode) (ans bool) {
	if root == nil {
		ans = true
		return
	}

	foundNil := false
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		q := []*TreeNode{}
		for _, node := range queue {
			if node == nil {
				foundNil = true
				continue
			}
			if foundNil { // 如果在同一层两次发现 nil, 则断定不是完全二叉树
				return false
			}
			q = append(q, node.Left)
			q = append(q, node.Right)
		}

		queue = q
	}

	ans = true
	return
}
