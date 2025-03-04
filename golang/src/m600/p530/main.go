package p530

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

const INTMAX = 1 << 31

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func preorderTraverse(node *TreeNode, ch chan<- *TreeNode) {
	if node == nil {
		return
	}
	preorderTraverse(node.Left, ch)
	ch <- node
	preorderTraverse(node.Right, ch)
}

func getMinimumDifference(root *TreeNode) int {
	var ch = make(chan *TreeNode, 0)

	go func() {
		preorderTraverse(root, ch)
		close(ch)
	}()

	var r = INTMAX
	prev := (*TreeNode)(nil)
	for node := range ch {
		if prev != nil {
			r = min(r, node.Val-prev.Val)
		}
		prev = node
	}
	return r
}
