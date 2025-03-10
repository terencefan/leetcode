package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val       int
	Neighbors []*Node
}

type TreeNodeNext struct {
	Val   int
	Left  *TreeNodeNext
	Right *TreeNodeNext
	Next  *TreeNodeNext
}

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}
