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

func lowestCommonAncestor2(root *TreeNode, p int, q int) (ans int) {
	curr := root
	for {
		ans = curr.Val
		if p < ans && q < ans {
			curr = curr.Left
		} else if p > ans && q > ans {
			curr = curr.Right
		} else {
			return
		}
	}
}
