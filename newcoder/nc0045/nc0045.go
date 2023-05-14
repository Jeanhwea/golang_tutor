package nc0045

// 滑动窗口的最大值
func maxInWindows(num []int, size int) (ans []int) {
	n := len(num)
	if size == 0 || size > n {
		return
	}
	var deque []int // 双端队列，维护单调递增的数值下标
	for curr := 0; curr < n; curr++ {
		// 如果当前值大于队尾，则队尾出队，直到当前值完全入队
		if len(deque) <= 0 {
			deque = append(deque, curr)
		} else {
			for len(deque) > 0 && num[deque[len(deque)-1]] < num[curr] {
				deque = deque[:len(deque)-1]
			}
			deque = append(deque, curr)
		}

		if curr < size-1 {
			continue
		}

		// 从队首查找到第一个不过期的下标即为当前的最大值
		var rear int
		for {
			rear = deque[0]
			if rear > curr-size {
				break
			}
			deque = deque[1:]
		}

		ans = append(ans, num[rear])
	}
	return
}
