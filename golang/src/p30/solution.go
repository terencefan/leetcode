package main

import (
	"fmt"
)

type TrieNode struct {
	arr      []*TrieNode
	terminal bool
	times    int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		make([]*TrieNode, 26),
		false,
		0,
	}
}

type Trie struct {
	root   *TrieNode
	times  int
	length int
}

func (t *Trie) Add(s string) {
	node := t.root
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if node.arr[c] == nil {
			node.arr[c] = NewTrieNode()
		}
		node = node.arr[c]
	}
	node.terminal = true
	node.times++
	t.times++
	t.length += len(s)
}

func add(node *TrieNode) {
	node.times++
}

func match(t *Trie, s string) bool {
	i, count := 0, 0

	node := t.root
	for i < t.length {
		c := s[i] - 'a'
		if node.arr[c] == nil {
			return false
		}
		node = node.arr[c]

		if node.terminal {
			if node.times == 0 {
				return false
			} else {
				node.times--
				count++
				defer add(node)
			}
			node = t.root
		}
		i++
	}

	return count == t.times
}

func findSubstring(s string, words []string) (r []int) {
	if s == "" || len(words) == 0 {
		return
	}

	t := &Trie{NewTrieNode(), 0, 0}
	for _, word := range words {
		t.Add(word)
	}

	for i := 0; i < len(s)-t.length+1; i++ {
		if match(t, s[i:]) {
			r = append(r, i)
		}
	}
	return
}

func main() {
	r := findSubstring("barfoothefoobarman", []string{"word", "good", "best", "good"})
	fmt.Println(r)
}
