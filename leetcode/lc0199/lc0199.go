package lc0199

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 二叉树的右视图
//
//    采用层序遍历, 每层将一层的最右边的数字放入 ans 结果中即可
//
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	que := []*TreeNode{root}
	for level := 0; len(que) > 0; level++ {
		prev := que
		que = nil
		for i, node := range prev {
			if i == 0 {
				ans = append(ans, node.Val)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
		}
	}

	return
}
