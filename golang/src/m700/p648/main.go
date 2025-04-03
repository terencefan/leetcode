package p648

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func (t *Trie) insert(word string) {
	node := t.root
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &TrieNode{}
		}
		node = node.children[c-'a']
	}
	node.word = word
}

func replaceWords(dictionary []string, sentence string) string {
	var tree = &Trie{root: &TrieNode{}}

	for _, word := range dictionary {
		tree.insert(word)
	}

	var r = make([]rune, 0)

	node, skip := tree.root, false
	for _, c := range sentence {
		if c == ' ' {
			r = append(r, ' ')
			skip = false
			node = tree.root
		} else if skip {
			continue
		} else if node != nil && len(node.word) > 0 {
			skip = true
		} else {
			r = append(r, c)
			if node != nil {
				node = node.children[c-'a']
			}
		}
	}
	return string(r)
}
