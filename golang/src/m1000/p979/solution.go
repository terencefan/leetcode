package main

var moves = 0

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func recursive(node *TreeNode) int {
	if node == nil {
		return 0
	}

	node.Val += recursive(node.Left)
	node.Val += recursive(node.Right)

	if node.Val == 1 {
		return 0
	} else {
		moves += abs(node.Val - 1)
		return node.Val - 1
	}
}

func distributeCoins(root *TreeNode) int {
	moves = 0
	recursive(root)
	return moves
}
