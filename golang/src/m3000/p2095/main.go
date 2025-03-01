package p2095

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func deleteMiddle(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head}
	var prev, node, node2 = dummy, head, head

	for node2 != nil && node2.Next != nil {
		prev = node
		node = node.Next
		node2 = node2.Next.Next
	}

	prev.Next = node.Next
	return dummy.Next
}
