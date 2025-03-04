package p92

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left >= right {
		return head
	}

	var dummy = &ListNode{Next: head}
	var prev, node = dummy, head

	if left > 0 {
		for range left - 1 {
			if node.Next == nil {
				return head
			}
			prev = node
			node = node.Next
		}
	}

	var reverseHead = prev
	var reverseTail = node

	for i := left; i < right; i++ {
		if node == nil {
			reverseHead.Next = prev
			break
		}
		next := node.Next
		node.Next, prev = prev, node
		node = next
	}

	reverseHead.Next = prev
	reverseTail.Next = node
	return dummy.Next
}
