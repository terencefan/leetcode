package main

import (
	"fmt"
)

type TrieNode struct {
	arr   []*TrieNode
	index int
	end   bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		arr:   make([]*TrieNode, 26),
		index: 0,
		end:   false,
	}
}

type Trie struct {
	root    *TrieNode
	indexes []int
}

func newTrie() *Trie {
	return &Trie{
		newTrieNode(),
		make([]int, 0),
	}
}

func testStr(word string, p, q int) bool {
	for p < q {
		if word[p] != word[q] {
			return false
		}
		p++
		q--
	}
	return true
}

func (t *Trie) add(index int, word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.arr[c] == nil {
			node.arr[c] = newTrieNode()
		}
		node = node.arr[c]
	}
	node.index = index
	node.end = true
}

func (t *Trie) match(right int, word string, full bool) []int {
	r := make([]int, 0)

	node := t.root
	for i := 0; i < len(word); i++ {
		if node.end == true && testStr(word, i, len(word)-1) {
			if node.index != right {
				r = append(r, node.index)
			}
		}

		c := word[i] - 'a'
		if node.arr[c] == nil {
			return r
		}
		node = node.arr[c]
	}

	if full {
		if node.end == true {
			if node.index != right {
				r = append(r, node.index)
			}
		}
	}
	return r
}

func reverse(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[len(s)-i-1] = s[i]
	}
	return string(r)
}

func palindromePairs(words []string) (r [][]int) {
	var t *Trie

	t = newTrie()
	for i, word := range words {
		t.add(i, word)
	}

	for i, word := range words {
		for _, index := range t.match(i, reverse(word), true) {
			r = append(r, []int{index, i})
		}
	}

	t = newTrie()
	for i, word := range words {
		t.add(i, reverse(word))
	}

	for i, word := range words {
		for _, index := range t.match(i, word, false) {
			r = append(r, []int{i, index})
		}
	}
	return r
}

func main() {
	r := palindromePairs([]string{"ab", "ba", "abc", "cba"})
	fmt.Println(r)
}
