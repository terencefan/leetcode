package p138

import "terencefan.com/leetcode/src/utils"

type Node = utils.RandomNode

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	var oldToNew = make(map[*Node]*Node)

	dummy := &Node{}
	node, prev := head, dummy
	for node != nil {
		newNode := &Node{Val: node.Val}
		oldToNew[node] = newNode
		prev.Next, prev = newNode, newNode
		node = node.Next
	}

	node = head
	newNode := dummy.Next
	for node != nil {
		newNode.Random = oldToNew[node.Random]
		node = node.Next
		newNode = newNode.Next
	}
	return dummy.Next
}
