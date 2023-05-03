package lc103

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 二叉树的锯齿形层序遍历
//
//   层序：搞一个队列来讲每层的节点记录
//   该题核心和层序遍历二叉树类似, 只是在奇数翻转一下数组
//
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	que := []*TreeNode{root}
	for level := 0; len(que) > 0; level++ {
		var vals []int
		prev := que
		que = nil
		for _, node := range prev {
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			vals = append(vals, node.Val)
		}

		if level % 2 == 1 {
			for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}

		ans = append(ans, vals)
	}

	return
}
