package p211

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func NewNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: NewNode()}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = NewNode()
		}
		node = node.children[c]
	}
	node.isWord = true
}

func search(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		return node.isWord
	}

	if word[index] == '.' {
		for _, child := range node.children {
			if search(child, word, index+1) {
				return true
			}
		}
	} else {
		if child, ok := node.children[rune(word[index])]; ok {
			return search(child, word, index+1)
		} else {
			return false
		}
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	return search(this.root, word, 0)
}
