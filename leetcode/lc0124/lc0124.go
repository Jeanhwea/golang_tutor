package lc0124

import (
	"math"

	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 二叉树中的最大路径和
//
//    记忆化递归搜索答案, memo
//      1. key:二叉树节点node
//      2. val:包含node的左右路径中较大的路径和, 在下面三者中取最大值
//         - 当前节点和左子树最大值
//         - 当前节点和右子树最大值
//         - 仅当前节点
//
//    算法: 先计算出 memo, 然后遍历 memo 求出最大路径和
//
func maxPathSum(root *TreeNode) (ans int) {
	memo = make(map[*TreeNode]int)
	calc(root)
	ans = math.MinInt
	for node := range memo {
		lsum := memo[node.Left]
		rsum := memo[node.Right]
		sum := node.Val
		if lsum > 0 {
			sum += lsum
		}
		if rsum > 0 {
			sum += rsum
		}
		if sum > ans {
			ans = sum
		}
	}
	return
}

var memo map[*TreeNode]int

func calc(node *TreeNode) (sum int) {
	if node == nil {
		sum = 0
		return
	}

	if val, ok := memo[node]; ok {
		sum = val
		return
	}

	max := 0
	lval, rval := calc(node.Left), calc(node.Right)
	if lval > max {
		max = lval
	}
	if rval > max {
		max = rval
	}

	sum = node.Val + max
	memo[node] = sum
	return
}
