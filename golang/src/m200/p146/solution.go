package main

import (
	"fmt"
)

type node struct {
	key, val  int
	pre, next *node
}

type LRUCache struct {
	head, tail *node
	cap, size  int
	m          map[int]*node
}

func Constructor(capacity int) LRUCache {
	head := &node{0, 0, nil, nil}
	tail := &node{0, 0, nil, nil}

	head.next = tail
	tail.pre = head

	return LRUCache{
		head, tail,
		capacity, 0,
		make(map[int]*node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.adjust(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) adjust(current *node) {
	current.pre.next = current.next
	current.next.pre = current.pre

	current.pre = this.head
	current.next = this.head.next
	this.head.next.pre = current
	this.head.next = current
}

func (this *LRUCache) Put(key int, value int) {
	if current, ok := this.m[key]; ok {
		current.val = value
		this.adjust(current)
	} else {
		if this.size == this.cap {
			t := this.tail.pre
			delete(this.m, t.key)
			t.pre.next = this.tail
			this.tail.pre = t.pre
			this.size--
		}
		current = &node{key, value, this.head, this.head.next}
		this.m[key] = current
		this.head.next.pre = current
		this.head.next = current
		this.size++
	}
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	fmt.Println(c)
}
