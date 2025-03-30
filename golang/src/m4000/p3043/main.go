package p3043

type TrieNode struct {
	Children [10]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(num int) {
	node := t.root

	n := int(1e8)
	for n > num {
		n /= 10
	}

	for n > 0 {
		d := num / n
		num %= n
		n /= 10

		if node.Children[d] == nil {
			node.Children[d] = &TrieNode{}
		}
		node = node.Children[d]
	}
}

func dfs(r *int, node1, node2 *TrieNode, k int) {
	*r = max(*r, k)
	for i := range 10 {
		if node1.Children[i] != nil && node2.Children[i] != nil {
			dfs(r, node1.Children[i], node2.Children[i], k+1)
		}
	}
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	var trie1, trie2 = Trie{&TrieNode{}}, Trie{&TrieNode{}}

	for _, num := range arr1 {
		trie1.insert(num)
	}

	for _, num := range arr2 {
		trie2.insert(num)
	}

	var r = 0
	dfs(&r, trie1.root, trie2.root, 0)
	return r
}
