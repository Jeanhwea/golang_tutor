package nc0100

import (
	"container/list"
)

// 双向链表 + hash 表 实现 LRU 缓存
type Entry struct {
	Key int
	Val int
}

type Solution struct {
	cap   int
	cache map[int]*list.Element
	lru   *list.List
}

func Constructor(capacity int) Solution {
	return Solution{
		cap:   capacity,
		cache: make(map[int]*list.Element, capacity),
		lru:   list.New(),
	}
}

func (s *Solution) get(key int) (ans int) {
	if node, ok := s.cache[key]; ok {
		if entry, ok := node.Value.(*Entry); ok {
			ans = entry.Val
		}
		s.lru.MoveToFront(node)
	} else {
		ans = -1
	}
	return
}

func (s *Solution) set(key int, value int) {
	if node, ok := s.cache[key]; ok {
		if entry, ok := node.Value.(*Entry); ok {
			entry.Val = value
		}
		s.lru.MoveToFront(node)
		return
	}

	if len(s.cache) >= s.cap {
		if e, ok := s.lru.Back().Value.(*Entry); ok {
			delete(s.cache, e.Key)
		}
		s.lru.Remove(s.lru.Back())
	}

	newEntry := &Entry{Key: key, Val: value}
	s.cache[key] = s.lru.PushFront(newEntry)
}
