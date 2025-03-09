package p148

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func sort(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	} else if head.Next == nil {
		return head, head
	}

	less, more := &ListNode{}, &ListNode{}

	l, m := less, more
	for node := head.Next; node != nil; node = node.Next {
		if node.Val < head.Val {
			l, l.Next = node, node
		} else {
			m, m.Next = node, node
		}
	}
	l.Next, m.Next = nil, nil

	lessHead, lessTail := sort(less.Next)
	moreHead, moreTail := sort(more.Next)

	var rhead, rtail *ListNode
	if lessHead == nil {
		rhead = head
	} else {
		rhead = lessHead
	}

	if moreTail == nil {
		rtail = head
	} else {
		rtail = moreTail
	}

	if lessTail != nil {
		lessTail.Next = head
	}
	head.Next = moreHead
	return rhead, rtail
}

func sortList1(head *ListNode) *ListNode {
	head, _ = sort(head)
	return head
}
