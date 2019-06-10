package main

import (
	"bytes"
	"container/heap"
	"fmt"
)

type TrieNode struct {
	nexts []*TrieNode
	count int
	str   string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		nexts: make([]*TrieNode, 27),
	}
}

func (n *TrieNode) get(c byte) *TrieNode {
	var i int
	if c == ' ' {
		i = 0
	} else {
		i = int(c - 'a' + 1)
	}

	if n.nexts[i] == nil {
		n.nexts[i] = NewTrieNode()
	}
	return n.nexts[i]
}

type Heap []*TrieNode

func (h Heap) Len() int {
	return len(h)
}

func greater(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return !greater(s2, s1)
	}
	for i := range s2 {
		if s1[i] > s2[i] {
			return true
		} else if s1[i] < s2[i] {
			return false
		}
	}
	return true
}

func (h Heap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return greater(h[i].str, h[j].str)
	} else {
		return h[i].count < h[j].count
	}
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*TrieNode))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

func (h Heap) peek() *TrieNode {
	return h[0]
}

func (n *TrieNode) dfs(h *Heap) {
	if n.str != "" {
		heap.Push(h, n)
		if h.Len() > 3 {
			heap.Pop(h)
		}
	}

	for i := 0; i < 27; i++ {
		if n.nexts[i] == nil {
			continue
		}
		n.nexts[i].dfs(h)
	}
}

func (n *TrieNode) candidates() (r []string) {
	h := make(Heap, 0, 4)
	n.dfs(&h)

	r = make([]string, h.Len())
	for h.Len() > 0 {
		r[h.Len()-1] = heap.Pop(&h).(*TrieNode).str
	}
	return r
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Add(s string, count int) {
	node := t.root
	for i := range s {
		node = node.get(s[i])
	}
	node.str = s
	node.count += count
}

type AutocompleteSystem struct {
	tree *Trie
	buf  *bytes.Buffer
	node *TrieNode
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	if len(times) != len(sentences) {
		panic("input data invalid!")
	}

	var tree = NewTrie()
	for i, sentence := range sentences {
		tree.Add(sentence, times[i])
	}

	return AutocompleteSystem{
		tree: tree,
		buf:  bytes.NewBuffer([]byte{}),
		node: tree.root,
	}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.tree.Add(string(this.buf.Bytes()), 1)
		this.buf.Reset()
		this.node = this.tree.root
		return []string{}
	}

	this.buf.WriteByte(c)
	this.node = this.node.get(c)
	return this.node.candidates()
}

func main() {
	system := Constructor(
		[]string{"abc", "abbc", "a"},
		[]int{3, 3, 3},
	)
	fmt.Println(system.Input('b'))
	fmt.Println(system.Input('c'))
	fmt.Println(system.Input('#'))
	fmt.Println(system.Input('b'))
	fmt.Println(system.Input('c'))
	fmt.Println(system.Input('#'))
	fmt.Println(system.Input('a'))
	fmt.Println(system.Input('b'))
	fmt.Println(system.Input('c'))
	fmt.Println(system.Input('#'))
	fmt.Println(system.Input('a'))
}
