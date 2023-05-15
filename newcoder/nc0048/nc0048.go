package nc0048

import (
	"container/heap"

	. "github.com/jeanhwea/golang_tutor/common/model"
)

var (
	smallPart *MaxHeap // 大顶堆, 存数据较小的部分
	largePart *MinHeap // 小顶堆, 存数据较大的部分

	// (smallPart ∪ largePart) 是所有数据流中的数据, counter 表示所有数据流计数
	counter int
)

func init() {
	smallPart = &MaxHeap{}
	largePart = &MinHeap{}
}

func Insert(num int) {
	counter++
	if counter%2 == 1 {
		heap.Push(largePart, num)
		minVal := heap.Pop(largePart)
		heap.Push(smallPart, minVal)
	} else {
		heap.Push(smallPart, num)
		maxVal := heap.Pop(smallPart)
		heap.Push(largePart, maxVal)
	}
}

func GetMedian() (medium float64) {
	if counter%2 == 1 {
		medium = float64(smallPart.Top())
	} else {
		medium = float64(smallPart.Top()+largePart.Top()) / 2.0
	}
	return
}
