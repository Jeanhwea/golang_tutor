package nc0100

// 双向链表 + hash 表 实现 LRU 缓存
type Entry struct {
	Key        int
	Val        int
	Prev, Next *Entry
}

type Solution struct {
	cap        int
	cache      map[int]*Entry
	head, tail *Entry
}

func Constructor(capacity int) Solution {
	dummy := &Entry{}
	return Solution{
		cap:   capacity,
		cache: make(map[int]*Entry, capacity),
		head:  dummy,
		tail:  dummy,
	}
}

func (s *Solution) get(key int) (ans int) {
	if node, ok := s.cache[key]; ok {
		ans = node.Val
		s.evict(node)
		s.pushFront(node)
	} else {
		ans = -1
	}
	return
}

func (s *Solution) set(key int, value int) {
	if node, ok := s.cache[key]; ok {
		node.Val = value
		s.evict(node)
		s.pushFront(node)
		return
	}

	newNode := &Entry{Key: key, Val: value}
	s.cache[key] = newNode
	if len(s.cache) <= s.cap {
		s.pushFront(newNode)
	} else {
		delete(s.cache, s.tail.Key)
		s.evict(s.tail)
		s.pushFront(newNode)
	}
}

func (s *Solution) evict(node *Entry) {
	if node == s.tail {
		node.Prev.Next, s.tail = node.Next, node.Prev
	} else {
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	}
	node.Next, node.Prev = nil, nil
}

func (s *Solution) pushFront(node *Entry) {
	if s.head == s.tail {
		s.head.Next, node.Prev = node, s.head
		s.tail = node
	} else {
		s.head.Next, node.Prev, node.Next, s.head.Next.Prev = node, s.head, s.head.Next, node
	}
}
