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
	cap      int                   // LFU 的最大容量
	minFreq  int                   // 当前最小的出现频率值
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

	if len(l.dataHash) >= l.cap {
		l.evict()
	}

	l.minFreq = 1
	l.write(key, value, l.minFreq)
}

// 驱逐淘汰的 key
func (l *LfuCache) evict() {
	lowerList := l.freqHash[l.minFreq]
	oldestKey := lowerList.Back().Value.(*Entry).Key

	lowerList.Remove(lowerList.Back())
	if lowerList.Len() == 0 {
		delete(l.freqHash, l.minFreq)
	}

	delete(l.dataHash, oldestKey)
}

// 双写到 hash 表中
func (l *LfuCache) write(key, value, freq int) {
	if _, ok := l.freqHash[freq]; !ok {
		l.freqHash[freq] = list.New()
	}

	if dest, ok := l.freqHash[freq]; ok {
		dest.PushFront(&Entry{Key: key, Val: value, Freq: freq})
		l.dataHash[key] = dest.Front()
	} else {
		panic("write empty list!")
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

	l.write(key, value, freq+1)
}
