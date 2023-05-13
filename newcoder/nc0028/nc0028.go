package nc0028

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func maxDepth(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
