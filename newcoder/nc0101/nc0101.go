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
	dataHash map[int]*list.Element // key: 键,   val: 元素节点
}

func NewLfuCache(capacity int) *LfuCache {
	return &LfuCache{
		cap:      capacity,
		minFreq:  0,
		freqHash: make(map[int]*list.List),
		dataHash: make(map[int]*list.Element),
	}
}

func (l *LfuCache) Get(key int) (value int) {
	value = -1

	if node, ok := l.dataHash[key]; ok {
		e := node.Value.(*Entry)
		value = e.Val
		l.update(node, key, value)
	}
	return
}

func (l *LfuCache) Set(key, value int) {
	if _, ok := l.dataHash[key]; ok {
		l.update(l.dataHash[key], key, value)
		return
	}

	// key 在 dataHash 里面不存在
	if len(l.dataHash) >= l.cap { // 如果到达容量, 需要驱逐最不常用的 key
		lowerList := l.freqHash[l.minFreq]
		oldestKey := lowerList.Back().Value.(*Entry).Key

		lowerList.Remove(lowerList.Back())
		if lowerList.Len() == 0 {
			delete(l.freqHash, l.minFreq)
		}

		delete(l.dataHash, oldestKey)
	}

	// 双写 key 到对应的 hash 表中
	l.minFreq = 1
	if _, ok := l.freqHash[1]; !ok {
		l.freqHash[1] = list.New()
	}

	if dest, ok := l.freqHash[1]; ok {
		dest.PushFront(&Entry{Key: key, Val: value, Freq: 1})
		l.dataHash[key] = dest.Front()
	}
}

func (l *LfuCache) update(ele *list.Element, key, value int) {
	freq := ele.Value.(*Entry).Freq

	if freqList, ok := l.freqHash[freq]; ok {
		freqList.Remove(ele)
		if freqList.Len() == 0 {
			delete(l.freqHash, freq)
			if l.minFreq == freq {
				l.minFreq++
			}
		}
	}

	if _, ok := l.freqHash[freq+1]; !ok {
		l.freqHash[freq+1] = list.New()
	}

	if dest, ok := l.freqHash[freq+1]; ok {
		dest.PushFront(&Entry{Key: key, Val: value, Freq: freq + 1})
		l.dataHash[key] = dest.Front()
	} else {
		panic("update: freq+1 has no list!")
	}
}
