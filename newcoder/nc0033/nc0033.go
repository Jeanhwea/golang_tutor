package nc0033

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func Mirror(root *TreeNode) (ans *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = Mirror(root.Right), Mirror(root.Left)

	ans = root
	return
}
