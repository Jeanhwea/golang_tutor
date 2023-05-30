package nc0101

// 设计LFU缓存结构
func LFU(operators [][]int, k int) (ans []int) {
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
