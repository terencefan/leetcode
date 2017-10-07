package p530

import (
	"ds"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTreeNode(v int) (node *TreeNode) {
	node = &TreeNode{}
	node.Val = v
	return
}

func walk(root *TreeNode) (t []int) {
	var s1 = ds.NewStack(root)
	var s2 = ds.NewStack()

	for {
		if s1.Len() > 0 {
			var node = s1.Pop().(*TreeNode)
			if node.Left != nil {
				s1.Push(node.Left)
			}
			s2.Push(node)
		} else if s2.Len() > 0 {
			var node = s2.Pop().(*TreeNode)
			if node.Right != nil {
				s1.Push(node.Right)
			}
			t = append(t, node.Val)
		} else {
			break
		}
	}
	return t
}

func GetMinimumDifference(root *TreeNode) (r int) {
	var t = walk(root)

	r = t[1] - t[0]
	for i := 2; i < len(t); i++ {
		if t[i]-t[i-1] < r {
			r = t[i] - t[i-1]
		}
	}
	return
}
