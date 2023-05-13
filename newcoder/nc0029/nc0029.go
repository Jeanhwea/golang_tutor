package nc0029

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func hasPathSum(root *TreeNode, sum int) bool {
	bfs(root, 0, sum)
	return found
}

var (
	found = false
)

func bfs(curr *TreeNode, parentSum, sum int) {
	if found || curr == nil {
		return
	}

	currSum := parentSum + curr.Val
	if curr.Left == nil && curr.Right == nil {
		if currSum == sum {
			found = true
		}
		return
	}

	if curr.Left != nil {
		bfs(curr.Left, currSum, sum)
	}

	if curr.Right != nil {
		bfs(curr.Right, currSum, sum)
	}
}
