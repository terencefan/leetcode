package p105

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]

	node := &TreeNode{Val: rootVal}

	k := 0
	for ; k < len(inorder); k++ {
		if inorder[k] == rootVal {
			break
		}
	}

	node.Left = buildTree(preorder[1:k+1], inorder[:k])
	node.Right = buildTree(preorder[k+1:], inorder[k+1:])
	return node
}
