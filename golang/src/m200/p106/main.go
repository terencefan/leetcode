package p106

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]

	node := &TreeNode{Val: rootVal}

	k := 0
	for ; k < len(inorder); k++ {
		if inorder[k] == rootVal {
			break
		}
	}

	node.Left = buildTree(inorder[:k], postorder[:k])
	node.Right = buildTree(inorder[k+1:], postorder[k:len(postorder)-1])
	return node
}
