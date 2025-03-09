package p139

type TrieNode struct {
	children []*TrieNode
	isWord   bool
}

func NewNode() *TrieNode {
	return &TrieNode{children: make([]*TrieNode, 26)}
}

func (t *TrieNode) Next(c byte) *TrieNode {
	return t.children[c-'a']
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
	node.isWord = true
}

func dfs(s string, index int, tree *Trie) bool {
	if index == len(s) {
		return true
	}

	node := tree.root
	stack := make([]int, 0)

	for i := index; i < len(s); i++ {
		node = node.Next(s[i])
		if node == nil {
			break
		} else if node.isWord {
			stack = append(stack, index+1)
		}
	}

	for len(stack) > 0 {
		idx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if dfs(s, idx, tree) {
			return true
		}
	}
	return false
}

func wordBreak2(s string, wordDict []string) bool {
	tree := &Trie{root: NewNode()}
	for _, word := range wordDict {
		tree.Insert(word)
	}
	return dfs(s, 0, tree)
}
