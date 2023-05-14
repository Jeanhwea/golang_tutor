package nc0041

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 层序遍历获取右视图
func solve(pre, vin []int) (ans []int) {
	root := cons(pre, vin)
	if root == nil {
		return
	}
	que := []*TreeNode{root}
	for level := 0; len(que) > 0; level++ {
		ans = append(ans, que[len(que)-1].Val)
		q := []*TreeNode{}
		for _, node := range que {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		que = q
	}
	return
}

// 重建二叉树
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
