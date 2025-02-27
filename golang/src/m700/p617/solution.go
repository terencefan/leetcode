package p617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil {
		return &TreeNode{t2.Val, t2.Left, t2.Right}
	} else if t2 == nil {
		return &TreeNode{t1.Val, t1.Left, t1.Right}
	} else {
		return &TreeNode{t1.Val + t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
	}
}
