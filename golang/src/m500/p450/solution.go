package main

func predecessor(node *TreeNode) *TreeNode {
	node = node.Left
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func successor(node *TreeNode) *TreeNode {
	node = node.Right
	if node == nil {
		return nil
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left != nil {
			pre := predecessor(root)
			root.Val = pre.Val
			root.Left = deleteNode(root.Left, root.Val)
		} else {
			post := successor(root)
			root.Val = post.Val
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}
