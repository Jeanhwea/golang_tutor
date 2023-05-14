package nc0040

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func reConstructBinaryTree(pre []int, vin []int) (root *TreeNode) {
	return cons(pre, vin)
}

func cons(pre, vin []int) (root *TreeNode) {
	if len(pre) == 0 {
		return
	}
	val, i, n := pre[0], 0, len(pre)
	for ; i < len(vin); i++ {
		if vin[i] == val {
			break
		}
	}
	root = &TreeNode{
		Val:   val,
		Left:  cons(pre[1:i+1], vin[0:i]),
		Right: cons(pre[i+1:n], vin[i+1:n]),
	}
	return
}
