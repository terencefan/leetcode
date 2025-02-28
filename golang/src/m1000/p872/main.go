package p872

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func iterate(root *TreeNode, ch chan<- int) {
	if root == nil {
		return
	}

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]

		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Left == nil && node.Right == nil {
			ch <- node.Val
		}
	}
	close(ch)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go iterate(root1, ch1)
	go iterate(root2, ch2)

	for expected := range ch1 {
		actual, ok := <-ch2
		if expected != actual || !ok {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}
