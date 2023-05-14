package nc0042

// 用两个栈实现队列
var (
	stack1, stack2 []int
)

func Push(node int) {
	stack1 = append(stack1, node)
	return
}

func Pop() (top int) {
	// stack2 是入栈序列, 如果 stack2 不空, 直接将数据出栈
	if len(stack2) > 0 {
		top, stack2 = stack2[len(stack2)-1], stack2[0:len(stack2)-1]
		return
	}

	// 将 stack1 全部搬运到 stack2 中, 然后等待出栈
	for len(stack1) > 0 {
		top, stack1 = stack1[len(stack1)-1], stack1[0:len(stack1)-1]
		stack2 = append(stack2, top)
	}

	top = Pop()
	return
}
