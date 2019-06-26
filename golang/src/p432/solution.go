package main

type ListNode struct {
	pre, next *ListNode
	key string
	count int
}

type AllOne struct {
	head, tail *ListNode
	m map[string]*ListNode
}


/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		head: nil,
		tail: nil,
		m: make(map[string]*ListNode),
	}
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
	if node, ok := this.m[key]; ok {
		node.count++

		for node.pre != nil && node.count > node.pre.count {

			pre, next := node.pre.pre, node.next

			if next != nil {
				next.pre = node.pre
			} else {
				this.tail = node.pre
			}

			if pre != nil {
				pre.next = node
			} else {
				this.head = node
			}

			node.pre.next = next
			node.pre.pre = node

			node.next = node.pre
			node.pre = pre

		}

	} else {
		node := &ListNode{this.tail, nil, key, 1}

		if this.tail != nil {
			this.tail.next = node
		}
		this.tail = node

		if this.head == nil {
			this.head = node
		}

		this.m[key] = node
	}
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
	if node, ok := this.m[key]; ok {
		if node.count == 1 {
			delete(this.m, key)

			pre, next := node.pre, node.next

			if pre != nil {
				pre.next = next
			} else {
				// it is the first node
				this.head = next
			}

			if next != nil {
				next.pre = pre
			} else {
				// it is the last node
				this.tail = pre
			}
		} else {
			node.count--

			for node.next != nil && node.count < node.next.count {
				pre, next := node.pre, node.next.next

				if next != nil {
					next.pre = node
				} else {
					this.tail = node
				}

				if pre != nil {
					pre.next = node.next
				} else {
					this.head = node.next
				}

				node.next.pre = pre
				node.next.next = node

				node.pre = node.next
				node.next = next
			}
		}
	} else {
		// nothing will happen.
	}
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.head == nil {
		return ""
	} else {
		return this.head.key
	}
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.tail == nil {
		return ""
	} else {
		return this.tail.key
	}
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */