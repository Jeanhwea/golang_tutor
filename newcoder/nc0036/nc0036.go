package nc0036

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func IsBalanced_Solution(root *TreeNode) bool {
	memo = make(map[*TreeNode]int)
	return isBalanced(root)
}

func isBalanced(root *TreeNode) (balanced bool) {
	if root == nil {
		balanced = true
		return
	}
	diff := abs(getHeight(root.Left) - getHeight(root.Right))
	if diff > 1 {
		balanced = false
		return
	}
	balanced = isBalanced(root.Left) && isBalanced(root.Right)
	return
}

var (
	memo map[*TreeNode]int
)

func getHeight(root *TreeNode) (height int) {
	if root == nil {
		return
	}
	if h, ok := memo[root]; ok {
		height = h
		return
	}
	height = max(getHeight(root.Left), getHeight(root.Right)) + 1
	memo[root] = height
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
