package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := 0, 0
	var node *ListNode

	node = headA
	for node != nil {
		node = node.Next
		a++
	}

	node = headB
	for node != nil {
		node = node.Next
		b++
	}

	if a > b {
		for i := 0; i < a-b; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < b-1; i++ {
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	l, r := head, head
	for l.Next != nil && r.Next != nil && r.Next.Next != nil {
		l = l.Next
		r = r.Next.Next
		if l == r {
			r = l.Next
			l.Next = nil
			return getIntersectionNode(head, r)
		}
	}
	return nil
}
