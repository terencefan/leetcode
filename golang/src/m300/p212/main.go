package p212

type TrieNode struct {
	children map[byte]*TrieNode
	count    int
	word     string
}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	t.root.count++
	node := t.root
	for i := range word {
		b := word[i]
		if _, ok := node.children[b]; !ok {
			node.children[b] = NewTrieNode()
		}
		node = node.children[b]
		node.count++
	}
	node.word = word
}

func (t *Trie) Remove(word string) {
	t.root.count--
	node := t.root
	for i := range word {
		b := word[i]
		node = node.children[b]
		if node == nil {
			return
		}
		node.count--
	}
	node.word = ""
}

func find(words *map[string]bool, board [][]byte, node *TrieNode, i, j int) {
	b := board[i][j]
	if next, ok := node.children[b]; ok {
		if len(next.word) > 0 && next.count > 0 {
			(*words)[next.word] = true
		}

		board[i][j] = '.'
		defer func() { board[i][j] = b }()

		for _, d := range directions {
			dx, dy := i+d[0], j+d[1]
			if dx < 0 || dx >= len(board) || dy < 0 || dy >= len(board[0]) {
				continue
			}
			find(words, board, next, dx, dy)
		}
	}
}

func findWords(board [][]byte, words []string) []string {
	tree := &Trie{root: NewTrieNode()}
	for _, word := range words {
		tree.Insert(word)
	}

	var r = make([]string, 0)
	for i, row := range board {
		for j := range row {
			var words = make(map[string]bool)
			find(&words, board, tree.root, i, j)
			for word := range words {
				tree.Remove(word)
				r = append(r, word)
			}
		}
	}
	return r
}
