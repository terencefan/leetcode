package main

import (
	"fmt"
)

type Node struct {
	key, val int
}

type Slot struct {
	nodes []*Node
}

func (s *Slot) add(key, val int) {
	for _, node := range s.nodes {
		if node.key == key {
			node.val = val
			return
		}
	}
	s.nodes = append(s.nodes, &Node{key, val})
}

func (s *Slot) find(key int) int {
	for _, node := range s.nodes {
		if node.key == key {
			return node.val
		}
	}
	return -1
}

func (s *Slot) remove(key int) {
	for _, node := range s.nodes {
		if node.key == key {
			node.val = -1
			return
		}
	}
}

type MyHashMap struct {
	slots []Slot
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		slots: make([]Slot, 997),
	}
}

func (this *MyHashMap) hash(key int) int {
	if key >= 0 {
		return key % 997
	} else {
		return ^key % 997
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.slots[this.hash(key)].add(key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.slots[this.hash(key)].find(key)
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.slots[this.hash(key)].remove(key)
}

func main() {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 2)
	fmt.Println(m.Get(1))
	fmt.Println(m.Get(3))
	m.Put(2, 1)
	fmt.Println(m.Get(2))
	m.Remove(2)
	fmt.Println(m.Get(2))
}
