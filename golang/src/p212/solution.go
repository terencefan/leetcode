package main

import (
	"fmt"
)

type Board [][]byte

func (b Board) height() int {
	return len(b)
}

func (b Board) width() int {
	return len(b[0])
}

type Trie struct {
	root *TrieNode
	r    []string
}

type TrieNode struct {
	val      byte
	word     string
	children []*TrieNode
}

func newTrieNode(b byte) *TrieNode {
	return &TrieNode{
		b,
		"",
		make([]*TrieNode, 26),
	}
}

func newTrie() *Trie {
	return &Trie{newTrieNode(0), make([]string, 0)}
}

func (t *Trie) add(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = newTrieNode(byte(c + 'a'))
		}
		node = node.children[c]
	}
	node.word = word
}

func clear(node *TrieNode) {
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			clear(node.children[i])
			node.children[i] = nil
			node.word = ""
		}
	}
}

func (t *Trie) has(word string) bool {
	node := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			return false
		}
		node = node.children[c]
	}
	return len(node.word) > 0
}

var trie *Trie

func dfs(board Board, x, y int, node *TrieNode) {
	b := board[x][y]
	if b == 0 {
		return
	}

	c := b - 'a'
	if node.children[c] == nil {
		return
	}
	node = node.children[c]

	if node.word != "" {
		trie.r = append(trie.r, node.word)
		node.word = ""
	}

	board[x][y] = 0
	if x > 0 {
		dfs(board, x - 1, y, node)
	}
	if y > 0 {
		dfs(board, x, y - 1, node)
	}
	if x < board.height() - 1 {
		dfs(board, x + 1, y, node)
	}
	if y < board.width() - 1 {
		dfs(board, x, y + 1, node)
	}
	board[x][y] = b

	return
}

func findWords(b Board, words []string) []string {
	trie = newTrie()
	for _, word := range words {
		trie.add(word)
	}

	for i := 0; i < b.height(); i++ {
		for j := 0; j < b.width(); j++ {
			dfs(b, i, j, trie.root)
		}
	}
	return trie.r
}

func main() {
	// r := findWords([][]byte{
	// 	{'o', 'a', 'a', 'n'},
	// 	{'e', 't', 'a', 'e'},
	// 	{'i', 'h', 'k', 'r'},
	// 	{'i', 'f', 'l', 'v'},
	// }, []string{"oath", "pea", "eat", "rain"})
	r := findWords([][]byte{
		{'a', 'a'},
	}, []string{"a"})
	fmt.Println(r)
}
