package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	r     *TreeNode
	depth int
)

func find(node, p, q *TreeNode, d int) (m, n bool) {
	if node == nil {
		return
	} else if node == p {
		m = true
	} else if node == q {
		n = true
	}
	m1, n1 := find(node.Left, p, q, depth+1)
	m2, n2 := find(node.Right, p, q, depth+1)
	m = m || m1 || m2
	n = n || n1 || n2
	if m && n && d > depth {
		r = node
		depth = d
	}
	return
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	r = root
	depth = 0
	find(root, p, q, 0)
	return r
}

func main() {

}
