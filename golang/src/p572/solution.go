package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return subtree(s, t)
}

func equal(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	} else {
		return s.Val == t.Val && equal(s.Left, t.Left) && equal(s.Right, t.Right)
	}
}

func subtree(s, t *TreeNode) bool {
	if t == nil {
		return true
	} else if s == nil {
		return false
	}
	return equal(s, t) || subtree(s.Left, t) || subtree(s.Right, t)
}
