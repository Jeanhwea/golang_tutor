package nc0039

import (
	"fmt"
	"strconv"

	. "github.com/jeanhwea/golang_tutor/common/model"
)

func Serialize(root *TreeNode) (s string) {
	if root == nil {
		s = "#"
		return
	}
	val := strconv.Itoa(root.Val)
	left, right := Serialize(root.Left), Serialize(root.Right)
	s = fmt.Sprintf("*%v,%v,%v", val, left, right)
	return
}

func Deserialize(s string) (root *TreeNode) {
	root, _ = deserialize(s, 0)
	return
}

func deserialize(s string, curr int) (root *TreeNode, next int) {
	n := len(s)
	if curr >= n {
		return
	}

	switch prelude := s[curr]; prelude {
	case '#':
		for ; curr < n; curr++ {
			if s[curr] == ',' {
				break
			}
		}
		next = curr + 1
		return
	case '*':
		prev := curr
		for ; curr < n; curr++ {
			if s[curr] == ',' {
				break
			}
		}
		val, _ := strconv.Atoi(s[prev+1 : curr])
		left, next1 := deserialize(s, curr+1)
		right, next2 := deserialize(s, next1)
		root = &TreeNode{
			Val:   val,
			Left:  left,
			Right: right,
		}
		next = next2
	}
	return
}
