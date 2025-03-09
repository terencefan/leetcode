package p30

const (
	D = iota
	T
	F
)

type TrieNode struct {
	children []*TrieNode
	count    int
	word     string
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
	node.word = word
	node.count++
}

func (t *Trie) Through(r *[]int, s string, index, length int) {
	var node = t.root
	var words = make([]string, 0)
	var left = index

	for i := index; i < len(s); i++ {
		node = node.Next(s[i])
		if node == nil {
			break
		} else if len(node.word) > 0 {
			for node.count < 1 && len(words) > 0 {
				word := words[0]
				words = words[1:]
				left += len(word)
				t.Insert(word)
			}
			if node.count > 0 {
				node.count--
				words = append(words, node.word)
				node = t.root
			}
			if len(words) == length {
				(*r)[left] = T
			}
		}
	}

	for _, word := range words {
		t.Insert(word)
	}
}

func (t *Trie) Run(s string, length int) []int {
	var r = make([]int, len(s))
	for i := range s {
		if t.root.Next(s[i]) == nil {
			r[i] = F
		} else if r[i] == D {
			t.Through(&r, s, i, length)
		}
	}

	res := make([]int, 0)
	for i, num := range r {
		if num == T {
			res = append(res, i)
		}
	}
	return res
}

func findSubstring(s string, words []string) []int {
	var tree = &Trie{root: NewNode()}

	for _, word := range words {
		tree.Insert(word)
	}
	return tree.Run(s, len(words))
}
