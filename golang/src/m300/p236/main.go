package p236

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func find(node, p, q *TreeNode) (*TreeNode, bool, bool) {
	if node == nil {
		return nil, false, false
	}

	fp, fq := false, false
	if node == p {
		fp = true
	} else if node == q {
		fq = true
	}

	left, lp, lq := find(node.Left, p, q)
	if lp && lq {
		return left, true, true
	}

	right, rp, rq := find(node.Right, p, q)
	if rp && rq {
		return right, true, true
	}
	return node, fp || lp || rp, fq || lq || rq
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node, _, _ := find(root, p, q)
	return node
}
