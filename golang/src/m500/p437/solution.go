package p437

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive1(root *TreeNode, sum int) (r int) {
	if root == nil {
		return
	} else if root.Val == sum {
		r++
	}
	r += recursive1(root.Left, sum-root.Val)
	r += recursive1(root.Right, sum-root.Val)
	fmt.Println(root.Val, sum, r)
	return
}

func recursive(root *TreeNode, sum int) (r int) {
	if root == nil {
		return
	}
	r += recursive(root.Left, sum) + pathSum(root.Right, sum)
	r += recursive1(root, sum)
	return r
}

func pathSum(root *TreeNode, sum int) int {
	r := recursive(root, sum)
	return r
}
