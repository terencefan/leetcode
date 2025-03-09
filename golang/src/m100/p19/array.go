package p19

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var nodes = []*ListNode{}

	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	m := len(nodes)
	if n >= m {
		return head.Next
	} else if n == 0 {
		nodes[m-n-1].Next = nil
	} else {
		nodes[m-n-1].Next = nodes[m-n].Next
	}
	return head
}
