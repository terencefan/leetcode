package main

import (
	"fmt"
)

type TrieNode struct {
	children[] *TrieNode
	count int
}

func (t *TrieNode) Add(word string, offset int) *TrieNode {
	t.count += 1
	if offset == len(word) {
		return t
	}
	c := int(word[offset]) - int('a')
	if t.children[c] == nil {
		t.children[c] = NewTrieNode()
	}
	return t.children[c].Add(word, offset + 1)
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
		count: 0,
	}
}

func makeAbbv(word string) string {
	if len(word) < 4 {
		return word
	}
	return fmt.Sprintf("%c%d%c", word[0], len(word) - 2, word[len(word) - 1])
}

func wordsAbbreviation(dict []string) []string {
	var m = make(map[string]interface{}, 0)

	for _, word := range dict {
		abbv := makeAbbv(word)
		if m[abbv] == nil {
			m[abbv] = word
			continue
		}

		switch v := m[abbv].(type) {
		case string:
			root := NewTrieNode()
			m[abbv] = root
			root.Add(v, 0)
			root.Add(word, 0)
		case *TrieNode:
			v.Add(word, 0)
		}
	}

	r := make([]string, 0)
	for _, word := range dict {
		abbv := makeAbbv(word)

		switch v := m[abbv].(type) {
		case string:
			r = append(r, abbv)
		case *TrieNode:
			offset := 0
			for v != nil && v.count > 1 {
				v = v.children[int(word[offset]) - int('a')]
				offset += 1
			}
			if offset < len(word) - 2 && offset > 0 {
				r = append(r, word[0:offset - 1] + makeAbbv(word[offset - 1:]))
			} else {
				r = append(r, word)
			}
		}

	}
	return r
}

func main() {
	r := wordsAbbreviation([]string{
		"abcdefg","abccefg","abcceft",
		// "like", "god", "internal", "me", "internet", "interval", "intension", "face", "intrusion",
	})
	fmt.Println(r)
}