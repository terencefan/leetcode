package p143

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	l, r := head, head

	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
	}

	prev, node := l, l.Next
	l.Next = nil

	for node != nil {
		next := node.Next
		node.Next = prev
		prev, node = node, next
	}

	n1, n2 := head, prev

	for n1 != nil && n2 != nil {
		n1n, n2n := n1.Next, n2.Next
		n1.Next = n2

		if n2 != n1n {
			n2.Next = n1n
		}
		n1, n2 = n1n, n2n
	}
}
