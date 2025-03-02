package p208

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func Constructor() Trie {
	return Trie{newTrieNode()}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = newTrieNode()
		}
		node = node.children[c]
	}
	node.isWord = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		node = node.children[c]
		if node == nil {
			return false
		}
	}
	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		node = node.children[c]
		if node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
