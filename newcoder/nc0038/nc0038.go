package nc0038

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 最近公共祖先
func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) (ans int) {
	node := dfs(root, o1, o2)
	ans = node.Val
	return
}

// 采用深度优先搜索
func dfs(root *TreeNode, p, q int) (ans *TreeNode) {
	if root == nil || root.Val == p || root.Val == q {
		return root
	}

	rootL, rootR := dfs(root.Left, p, q), dfs(root.Right, p, q)

	// 情况一: 说明 p, q 分布在 root 的左右子树, root 是最近公共祖先
	if rootL != nil && rootR != nil {
		ans = root
		return
	}

	// 情况二: p, q 都在 root 的右子树, 返回右子树搜索的最近公共祖先结果 rootR
	if rootL == nil {
		ans = rootR
		return
	}

	// 情况三: p, q 都在 root 的左子树, 返回左子树搜索的最近公共祖先结果 rootL
	if rootR == nil {
		ans = rootL
		return
	}

	// 不存在 rootL == nil && rootR == nil
	return
}
