package p19

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{Next: head}
	node1, node2 := head, head

	for range n {
		node2 = node2.Next
	}

	prev := dummy
	for node2 != nil {
		prev = node1
		node1 = node1.Next
		node2 = node2.Next
	}
	if n == 0 {
		prev.Next = nil
	} else {
		prev.Next = node1.Next
	}
	return dummy.Next
}
