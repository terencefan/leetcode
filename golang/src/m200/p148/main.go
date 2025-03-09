package p148

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n1, n2 := head, head

	for n2.Next != nil && n2.Next.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
	}

	p1 := sortList(n1.Next)
	n1.Next = nil
	p2 := sortList(head)

	dummy := &ListNode{}
	prev := dummy
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			prev.Next = p1
			p1 = p1.Next
		} else {
			prev.Next = p2
			p2 = p2.Next
		}
		prev = prev.Next
	}

	if p1 != nil {
		prev.Next = p1
	} else if p2 != nil {
		prev.Next = p2
	}
	return dummy.Next
}
