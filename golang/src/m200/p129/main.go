package p129

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

type NodeVal struct {
	val  int
	node *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	r := 0
	q := []NodeVal{{val: root.Val, node: root}}
	for len(q) > 0 {
		l := len(q)

		for i := range l {
			nv := q[i]
			if nv.node.Left == nil && nv.node.Right == nil {
				r += nv.val
				continue
			}
			if nv.node.Left != nil {
				q = append(q, NodeVal{val: nv.val*10 + nv.node.Left.Val, node: nv.node.Left})
			}
			if nv.node.Right != nil {
				q = append(q, NodeVal{val: nv.val*10 + nv.node.Right.Val, node: nv.node.Right})
			}
		}
		q = q[l:]
	}
	return r
}
