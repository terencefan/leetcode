package main

var m int

const (
	INTMIN = -(1 << 32)
	INTMAX = 1<<32 - 1
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func isBST(node *TreeNode) (size, lower, upper int) {
	if node.Left == nil && node.Right == nil {
		m = max(m, 1)
		return 1, node.Val, node.Val
	} else if node.Left == nil {
		size, lower, upper := isBST(node.Right)
		if node.Val < lower {
			m = max(m, size+1)
			return size + 1, node.Val, upper
		} else {
			return 0, INTMIN, INTMAX
		}
	} else if node.Right == nil {
		size, lower, upper := isBST(node.Left)
		if node.Val > upper {
			m = max(m, size+1)
			return size + 1, lower, node.Val
		} else {
			return 0, INTMIN, INTMAX
		}
	} else {
		s1, l1, u1 := isBST(node.Left)
		s2, l2, u2 := isBST(node.Right)
		if node.Val > u1 && node.Val < l2 {
			m = max(m, s1+s2+1)
			return s1 + s2 + 1, l1, u2
		} else {
			return 0, INTMIN, INTMAX
		}
	}
}

func largestBSTSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	m = 0
	isBST(root)
	return m
}
