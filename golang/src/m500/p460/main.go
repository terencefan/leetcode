package p460

type listNode[T any] struct {
	value      T
	prev, next *listNode[T]
}

type LinkedSet[T comparable] struct {
	m          map[T]*listNode[T]
	head, tail *listNode[T]
}

func NewLinkedList[T comparable]() *LinkedSet[T] {
	return &LinkedSet[T]{
		m: make(map[T]*listNode[T]),
	}
}

func (s *LinkedSet[T]) Add(val T) {
	if _, ok := s.m[val]; ok {
		return
	}

	node := &listNode[T]{value: val}
	s.m[val] = node

	if s.head == nil {
		s.head, s.tail = node, node
		return
	} else {
		s.tail.next = node
		node.prev = s.tail
		s.tail = node
	}
}

func (s *LinkedSet[T]) Count() int {
	return len(s.m)
}

func (s *LinkedSet[T]) Remove(val T) {
	if node, ok := s.m[val]; !ok {
		return
	} else {
		delete(s.m, val)
		if node.prev == nil && node.next == nil {
			s.head, s.tail = nil, nil
		} else if node.prev == nil {
			s.head = node.next
			s.head.prev = nil
		} else if node.next == nil {
			s.tail = node.prev
			s.tail.next = nil
		} else {
			node.prev.next = node.next
			node.next.prev = node.prev
		}
	}
}

func (s *LinkedSet[T]) First() T {
	return s.head.value
}

type ValueCount struct {
	val, count int
}

type LFUCache struct {
	valuesMap  map[int]*ValueCount
	keySetsMap map[int]*LinkedSet[int]
	capacity   int
	minCount   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		valuesMap:  make(map[int]*ValueCount),
		keySetsMap: make(map[int]*LinkedSet[int]),
		capacity:   capacity,
		minCount:   0,
	}
}

func (c *LFUCache) Get(key int) int {
	if item, ok := c.valuesMap[key]; ok {
		c.keySetsMap[item.count].Remove(key)

		if c.minCount == item.count && c.keySetsMap[item.count].Count() == 0 {
			c.minCount++
		}

		item.count++
		if c.keySetsMap[item.count] == nil {
			c.keySetsMap[item.count] = NewLinkedList[int]()
		}
		c.keySetsMap[item.count].Add(key)
		return item.val
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if item, ok := c.valuesMap[key]; ok {
		c.keySetsMap[item.count].Remove(key)

		if c.minCount == item.count && c.keySetsMap[item.count].Count() == 0 {
			c.minCount++
		}

		item.count++
		item.val = value
		if c.keySetsMap[item.count] == nil {
			c.keySetsMap[item.count] = NewLinkedList[int]()
		}
		c.keySetsMap[item.count].Add(key)
		return
	} else if len(c.valuesMap) == c.capacity {
		s := c.keySetsMap[c.minCount]
		delete(c.valuesMap, s.First())
		s.Remove(s.First())
	}

	c.minCount = 1
	c.valuesMap[key] = &ValueCount{val: value, count: 1}
	if c.keySetsMap[1] == nil {
		c.keySetsMap[1] = NewLinkedList[int]()
	}
	c.keySetsMap[1].Add(key)
}
