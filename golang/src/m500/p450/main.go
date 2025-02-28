package p450

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func next(node *TreeNode) *TreeNode {
	if node.Right == nil {
		return nil
	}

	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func prev(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return nil
	}

	node = node.Left
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func search(parent *TreeNode, node *TreeNode, key int) (p1 *TreeNode, p2 *TreeNode) {
	if node == nil {
		return parent, nil
	} else if node.Val < key {
		return search(node, node.Right, key)
	} else if node.Val > key {
		return search(node, node.Left, key)
	} else {
		return parent, node
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	parent, node := search(nil, root, key)
	if node == nil {
		return root
	}

	prev, next := prev(node), next(node)

	if node.Left == nil && node.Right == nil {
		if parent == nil {
			return nil
		} else if parent.Left != nil && parent.Left.Val == key {
			parent.Left = nil
		} else if parent.Right != nil && parent.Right.Val == key {
			parent.Right = nil
		}
	} else if next != nil {
		var val = next.Val
		deleteNode(node, val)
		node.Val = val
	} else if prev != nil {
		var val = prev.Val
		deleteNode(node, val)
		node.Val = val
	}
	return root
}
