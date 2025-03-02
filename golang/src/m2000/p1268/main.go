package p1268

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(s string) {
	node := t.root
	for _, c := range s {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = NewTrieNode()
		}
		node = node.children[c-'a']
	}
	node.isWord = true
}

type TrieNode struct {
	children []*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make([]*TrieNode, 26), isWord: false}
}

func (t *TrieNode) next(c rune) *TrieNode {
	return t.children[c-'a']
}

func dfs(node *TrieNode, k *int, runes *[]rune, r *[]string) {
	if *k == 0 {
		return
	}
	if node.isWord {
		*r = append(*r, string(*runes))
		*k--
	}

	for i, next := range node.children {
		if next != nil {
			*runes = append(*runes, rune('a'+i))
			dfs(next, k, runes, r)
			*runes = (*runes)[:len(*runes)-1]
		}
	}
}

func (t *TrieNode) find(k int) []string {
	var runes = make([]rune, 0)
	var r = make([]string, 0)
	dfs(t, &k, &runes, &r)
	return r
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := Trie{root: NewTrieNode()}
	for _, product := range products {
		trie.insert(product)
	}

	var node = trie.root
	var r = make([][]string, len(searchWord))
	for i := range len(searchWord) {
		r[i] = []string{}
	}

	for i, c := range searchWord {
		node = node.next(c)
		if node == nil {
			break
		}
		for _, postfix := range node.find(3) {
			r[i] = append(r[i], searchWord[:i+1]+postfix)
		}
	}
	return r
}
