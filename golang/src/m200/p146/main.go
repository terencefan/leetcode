package p146

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	count      int
	capacity   int
	nodes      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		nodes:    make(map[int]*Node),
	}
}

func (this *LRUCache) moveToTail(node *Node) {
	if node == this.tail {
		return
	} else if node.prev != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	} else {
		this.head = node.next
		node.next.prev = nil
	}
	this.tail.next = node
	node.prev = this.tail
	node.next = nil
	this.tail = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodes[key]; ok {
		this.moveToTail(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodes[key]; ok {
		node.val = value
		this.moveToTail(node)
		return
	}
	if this.count == 0 {
		this.head = &Node{key: key, val: value}
		this.tail = this.head
		this.nodes[key] = this.tail
		this.count++
	} else {
		this.tail.next = &Node{key: key, val: value, prev: this.tail}
		this.tail = this.tail.next
		this.nodes[key] = this.tail
		this.count++

		if this.count > this.capacity {
			delete(this.nodes, this.head.key)
			this.head = this.head.next
			this.head.prev = nil
			this.count--
		}
	}
}
