package main

func resursive(n1, n2 *ListNode) int {
	var mod int
	if n1.Next != nil {
		mod = resursive(n1.Next, n2.Next)
	}
	n1.Val += n2.Val + mod
	if n1.Val > 9 {
		n1.Val -= 10
		return 1
	} else {
		return 0
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	n1, n2 := l1, l2
	for n1 != nil && n2 != nil {
		n1 = n1.Next
		n2 = n2.Next
	}

	if n1 == nil {
		n1, n2 = n2, n1
		l1, l2 = l2, l1
	}

	r1, r2 := &ListNode{0, l1}, &ListNode{0, nil}

	t := r2
	for n1 != nil {
		n1 = n1.Next
		t.Next = &ListNode{0, nil}
		t = t.Next
	}
	t.Next = l2

	resursive(r1, r2)
	if r1.Val == 0 {
		return r1.Next
	} else {
		return r1
	}
}
