package p2352

type TrieNode struct {
	children map[int]*TrieNode
	count    int
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(nums []int) {
	var node = t.root
	for _, num := range nums {
		if _, ok := node.children[num]; !ok {
			node.children[num] = &TrieNode{make(map[int]*TrieNode), 0}
		}
		node = node.children[num]
		node.count++
	}
}

func equalPairs(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	n := len(grid)

	var trie = Trie{root: &TrieNode{make(map[int]*TrieNode), 0}}
	for i := range n {
		trie.Insert(grid[i])
	}

	var r = 0
	var ok bool
	for i := range n {
		node := trie.root
		for j := range n {
			if node == nil {
				break
			}
			if node, ok = node.children[grid[j][i]]; !ok {
				break
			}
		}
		if node != nil {
			r += node.count
		}
	}
	return r
}
