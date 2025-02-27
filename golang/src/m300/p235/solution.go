package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p.Val > q.Val {
		p, q = q, p
	}

	node := root
	for node != nil {
		if p.Val == node.Val || q.Val == node.Val {
			return node
		} else if p.Val < node.Val && q.Val > node.Val {
			return node
		} else if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
