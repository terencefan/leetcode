package p141

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p1, p2 := head, head.Next

	for p2 != nil && p2.Next != nil {
		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return false
}
