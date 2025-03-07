package p138

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var node, res *Node

	node = head
	for node != nil {
		node.Next = &Node{
			Val:    node.Val,
			Next:   node.Next,
			Random: nil,
		}
		node = node.Next.Next
	}

	node = head
	for node != nil {
		real := node.Next
		if node.Random != nil {
			real.Random = node.Random.Next
		}
		node = real.Next
	}

	node, res = head, head.Next
	for node != nil {
		real := node.Next
		node.Next = real.Next
		node = node.Next

		if node != nil {
			real.Next = node.Next
		}
	}
	return res
}
