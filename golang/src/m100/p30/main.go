package p30

type TrieNode struct {
	children []*TrieNode
	count    int
}

func NewNode() *TrieNode {
	return &TrieNode{children: make([]*TrieNode, 26)}
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = NewNode()
		}
		node = node.children[c-'a']
	}
	node.count++
}

func (t *Trie) Run(s string, length int) []int {
	return []int{}
}

func findSubstring(s string, words []string) []int {
	var tree = &Trie{root: NewNode()}

	for _, word := range words {
		tree.Insert(word)
	}
	return tree.Run(s, len(words))
}
