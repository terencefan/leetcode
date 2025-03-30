package p24

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func swapPairs(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head}

	prev, node := dummy, head

	for node != nil && node.Next != nil {
		nnext := node.Next.Next
		prev.Next = node.Next
		node.Next.Next = node
		node.Next = nnext
		prev, node = node, nnext
	}
	return dummy.Next
}
