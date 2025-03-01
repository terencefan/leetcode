package p328

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var oddHead = head.Next

	var prev1, prev2 *ListNode

	var even = false
	for node := head; node != nil; node = node.Next {
		if prev2 != nil {
			prev2.Next = node
		}
		prev2 = prev1
		prev1 = node
		even = !even
	}
	if prev2 != nil {
		prev2.Next = nil
	}

	if even {
		if prev1 != nil {
			prev1.Next = oddHead
		}
	} else {
		if prev2 != nil {
			prev2.Next = oddHead
		}
	}
	return head
}
