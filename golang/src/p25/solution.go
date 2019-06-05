package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	root := &ListNode{-1, nil}
	first := root

	node := head

	var pre, left, right *ListNode

	for node != nil {
		left = node
		for i := 0; i < k; i++ {
			if node == nil {
				first.Next = left
				return root.Next
			}
			node = node.Next
			right = node
		}

		pre, node = nil, left
		for node != right {
			next := node.Next
			node.Next = pre
			pre = node
			node = next
		}
		first.Next = pre
		first = left
	}
	return root.Next
}
