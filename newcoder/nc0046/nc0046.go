package nc0046

// 最小的 K 个数
func GetLeastNumbers_Solution(input []int, k int) (ans []int) {
	var heap IntHeap
	for i, num := range input {
		heap.push(num)
		if i > k-1 {
			heap.pop()
		}
	}

	for !heap.empty() {
		ans = append(ans, heap.pop())
	}
	return
}

type IntHeap []int

func (h *IntHeap) empty() bool {
	return len(*h) == 0
}

func (h *IntHeap) push(num int) {
	*h = append(*h, num)
	h.swim()
}

func (h *IntHeap) pop() (top int) {
	n := len(*h)
	if n <= 0 {
		return
	}
	top = (*h)[0]
	(*h)[0], *h = (*h)[n-1], (*h)[:n-1]
	h.sink()
	return
}

// 上浮操作
func (h *IntHeap) swim() {
	n := len(*h)
	if n <= 0 {
		return
	}

	s := n - 1
	for {
		p := (s - 1) / 2
		if (*h)[p] >= (*h)[s] {
			break
		}
		(*h)[p], (*h)[s] = (*h)[s], (*h)[p]
		s = p
	}
}

// 下沉操作
func (h *IntHeap) sink() {
	n := len(*h)
	if n <= 0 {
		return
	}

	p := 0
	for {
		max := p
		s := p*2 + 1
		if s < n && (*h)[s] > (*h)[max] {
			max = s
		}
		s++
		if s < n && (*h)[s] > (*h)[max] {
			max = s
		}

		if max == p {
			break
		}

		(*h)[p], (*h)[max] = (*h)[max], (*h)[p]
		p = max
	}
}
