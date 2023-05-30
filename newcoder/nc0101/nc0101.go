package nc0101

// 设计LFU缓存结构
func LFU(operators [][]int, k int) (ans []int) {
	n := len(operators)
	lfu := NewLFU(k)
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

// 思路: 采用双 hash 表来记录数据
var (
	freqHash map[int]*Entry // key: 频率, val: 链表头
	nodeHash map[int]*Entry // key: 键,   val: 元素节点
)

type Entry struct {
	Key  int    // 键
	Val  int    // 值
	Freq int    // 频率
	Next *Entry // 节点下一个指针
}

type LFUCache struct {
}

func NewLFU(capacity int) LFUCache {
	return LFUCache{}
}

func (l *LFUCache) Get(key int) (value int) {
	return
}

func (l *LFUCache) Set(key, value int) {

}
