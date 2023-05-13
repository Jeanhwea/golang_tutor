package nc0036

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func IsBalanced_Solution(root *TreeNode) (ans bool) {
	_, ans = isBalanced(root)
	return
}

func isBalanced(root *TreeNode) (height int, balanced bool) {
	if root == nil {
		balanced = true
		return
	}
	heightL, ok := isBalanced(root.Left)
	if !ok {
		balanced = false
		return
	}
	heightR, ok := isBalanced(root.Right)
	if !ok {
		balanced = false
		return
	}
	diff := heightL - heightR
	if diff > 1 || diff < -1 {
		balanced = false
		return
	}

	height = max(heightL, heightR) + 1
	balanced = true
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
