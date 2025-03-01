package p1448

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func dfs(node *TreeNode, currentMax int) int {
	if node == nil {
		return 0
	} else if node.Val < currentMax {
		return dfs(node.Left, currentMax) + dfs(node.Right, currentMax)
	} else {
		currentMax = node.Val
		return dfs(node.Left, currentMax) + dfs(node.Right, currentMax) + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stack(root *TreeNode) int {
	var s1 = make([]*TreeNode, 0)
	var s2 = make([]int, 0)

	s1 = append(s1, root)
	s2 = append(s2, root.Val)
	var r = 0

	for len(s1) > 0 {
		node, currentMax := s1[len(s1)-1], s2[len(s2)-1]
		s1, s2 = s1[:len(s1)-1], s2[:len(s2)-1]

		if node == nil {
			continue
		} else if node.Val <= currentMax {
			r += 1
		}
		nextMax := max(currentMax, node.Val)

		if node.Left != nil {
			s1 = append(s1, node.Left)
			s2 = append(s2, nextMax)
		}
		if node.Right != nil {
			s1 = append(s1, node.Right)
			s2 = append(s2, nextMax)
		}
	}
	return r
}

func goodNodes(root *TreeNode) int {
	return stack(root)
}

func init() {
	goodNodes(nil)
}
