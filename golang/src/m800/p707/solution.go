package main

import (
	"fmt"
)

type MyLinkedNode struct {
	prev, next *MyLinkedNode
	val        int
}

func NewLinkedNode(val int) *MyLinkedNode {
	return &MyLinkedNode{nil, nil, val}
}

type MyLinkedList struct {
	head, tail *MyLinkedNode
	length     int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{nil, nil, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if node := this.index(index); node != nil {
		return node.val
	} else {
		return -1
	}
}

func (this *MyLinkedList) index(index int) *MyLinkedNode {
	if this.length == 0 {
		return nil
	} else if index < 0 {
		return nil
	}

	node := this.head
	for node.next != nil && index > 0 {
		node = node.next
		index--
	}

	if index == 0 {
		return node
	} else {
		return nil
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := NewLinkedNode(val)
	if this.length == 0 {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := NewLinkedNode(val)
	if this.length == 0 {
		this.head = node
		this.tail = node
	} else {
		node.prev = this.tail
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val) // which is fucking meaningless
	} else if index == 0 {
		this.AddAtHead(val)
	} else if index == this.length {
		this.AddAtTail(val)
	} else if index > this.length {
		return
	} else {
		next := this.index(index)
		node := NewLinkedNode(val)

		node.next = next
		node.prev = next.prev

		next.prev.next = node
		next.prev = node

		this.length++
	}
}

func (this *MyLinkedList) Len() int {
	return this.length
}

func (this *MyLinkedList) clear() {
	this.head = nil
	this.tail = nil
	this.length = 0
}

func (this *MyLinkedList) removeHead() {
	this.head = this.head.next
	this.head.prev = nil
	this.length--
}

func (this *MyLinkedList) removeTail() {
	this.tail = this.tail.prev
	this.tail.next = nil
	this.length--
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	} else if this.length == 0 {
		return
	} else if this.length == 1 {
		if index > 0 {
			return
		} else {
			this.clear()
		}
	} else if index >= this.Len() {
		return
	} else if index == 0 {
		this.removeHead()
	} else if index == this.Len()-1 {
		this.removeTail()
	} else {
		node := this.index(index)
		node.prev.next = node.next
		node.next.prev = node.prev
		this.length--
	}
}

func main() {
	obj := Constructor();
	fmt.Println(obj.Get(0))
	obj.AddAtIndex(1, 2);
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1))
	obj.AddAtIndex(0, 1);
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1))
}
