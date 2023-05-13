package nc0027

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func Print(root *TreeNode) (ans [][]int) {
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
		if level%2 == 1 {
			for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}
		ans = append(ans, vals)
		queue = newQueue
	}
	return
}
