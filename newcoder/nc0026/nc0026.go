package nc0026

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		newQueue := []*TreeNode{}
		var vals []int
		for _, node := range queue {
			vals = append(vals, node.Val)
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		ans = append(ans, vals)
		queue = newQueue
	}
	return
}
