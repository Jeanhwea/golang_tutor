package nc0101

import (
	"container/list"
)

// 设计 LFU 缓存结构
func LFU(operators [][]int, k int) (ans []int) {
	n := len(operators)
	lfu := NewLfuCache(k)
	for i := 0; i < n; i++ {
		switch opt := operators[i][0]; opt {
		case 1:
			lfu.Set(operators[i][1], operators[i][2])
		case 2:
			ans = append(ans, lfu.Get(operators[i][1]))
		}
	}
	return
}

type Entry struct {
	Key  int // 键
	Val  int // 值
	Freq int // 频率
}

// 思路: 采用双 hash 表来记录数据
type LfuCache struct {
	cap      int
	minFreq  int
	freqHash map[int]*list.List    // key: 频率, val: 链表头
	nodeHash map[int]*list.Element // key: 键,   val: 元素节点
}

func NewLfuCache(capacity int) *LfuCache {
	return &LfuCache{
		cap:      capacity,
		minFreq:  0,
		freqHash: make(map[int]*list.List),
		nodeHash: make(map[int]*list.Element),
	}
}

func (l *LfuCache) Get(key int) (value int) {
	if node, ok := l.nodeHash[key]; ok {
		e := node.Value.(*Entry)
		value = e.Val
		l.update(node)
	} else {
		value = -1
	}
	return
}

func (l *LfuCache) Set(key, value int) {
	if node, ok := l.nodeHash[key]; ok {
		e := node.Value.(*Entry)
		e.Val = value
		l.update(node)
	} else {
		newEntry := &Entry{Key: key, Val: value, Freq: 1}
		if newEntry.Freq >= l.minFreq {
			if h, ok1 := l.freqHash[newEntry.Freq]; ok1 {
				l.nodeHash[key] = h.PushBack(newEntry)
			} else {
				newList := list.New()
				l.nodeHash[key] = newList.PushBack(newEntry)
				l.freqHash[newEntry.Freq] = newList
			}
		}
	}
}

func (l *LfuCache) update(ele *list.Element) {
	if ele == nil {
		return
	}

	e := ele.Value.(*Entry)
	if head, ok := l.freqHash[e.Freq]; ok {
		// 将元素从旧的列表中移除
		if head.Len() == 1 {
			delete(l.freqHash, e.Freq)
		} else {
			head.Remove(ele)
		}

		// 尝试添加元素到新的列表中
		e.Freq++
		if newHead, found := l.freqHash[e.Freq]; found {
			newHead.PushBack(e)
		} else {
			newList := list.New()
			newList.PushBack(e)
			l.freqHash[e.Freq] = newList
		}
	} else {
		panic("element not found!")
	}

	// 驱逐多余的元素
	if len(l.nodeHash) > l.cap {
		for {
			if h, ok := l.freqHash[l.minFreq]; ok {
				oldest := h.Front().Value.(*Entry)
				delete(l.nodeHash, oldest.Key)
				if h.Len() == 1 {
					delete(l.freqHash, l.minFreq)
					l.minFreq++
				} else {
					h.Remove(h.Front())
				}
				break
			}
			l.minFreq++
		}
	}
}
