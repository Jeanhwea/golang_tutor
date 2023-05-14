package nc0038

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) (ans int) {
	node := dfs(root, o1, o2)
	ans = node.Val
	return
}

func dfs(root *TreeNode, p, q int) (ans *TreeNode) {
	if root == nil || root.Val == p || root.Val == q {
		return root
	}

	rootL, rootR := dfs(root.Left, p, q), dfs(root.Right, p, q)
	if rootL != nil && rootR != nil {
		ans = root
		return
	}
	if rootL == nil {
		ans = rootR
		return
	}
	if rootR == nil {
		ans = rootL
		return
	}
	return
}
