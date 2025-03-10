package p173

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var r = &BSTIterator{stack: []*TreeNode{}}
	r.push(root)
	return *r
}

func (this *BSTIterator) push(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.push(node.Right)
	return node.Val

}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
