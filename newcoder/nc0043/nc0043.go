package nc0043

var (
	stack, minStack []int
)

func Push(node int) {
	stack = append(stack, node)
	currMin := node
	if len(minStack) > 0 && minStack[len(minStack)-1] < currMin {
		currMin = minStack[len(minStack)-1]
	}
	minStack = append(minStack, currMin)
}

func Pop() {
	stack = stack[:len(stack)-1]
	minStack = minStack[:len(minStack)-1]
}

func Top() (top int) {
	top = stack[len(stack)-1]
	return
}

func Min() (minVal int) {
	minVal = minStack[len(minStack)-1]
	return
}
