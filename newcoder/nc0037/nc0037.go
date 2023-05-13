package nc0037

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func lowestCommonAncestor(root *TreeNode, p int, q int) (ans int) {
	ans = root.Val
	if ans > q && ans > p {
		ans = lowestCommonAncestor(root.Left, p, q)
		return
	}
	if ans < q && ans < p {
		ans = lowestCommonAncestor(root.Right, p, q)
		return
	}
	return
}
