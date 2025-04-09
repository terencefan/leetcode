package p2

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func iter(node *ListNode) <-chan *ListNode {
	ch := make(chan *ListNode, 10)

	go func() {
		for node != nil {
			ch <- node
			node = node.Next
		}
		close(ch)
	}()

	return ch
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	node := dummy

	carry := 0
	ch1, ch2 := iter(l1), iter(l2)

	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2

		if !ok1 && !ok2 {
			break
		}

		sum := carry
		if ok1 {
			sum += n1.Val
		}
		if ok2 {
			sum += n2.Val
		}

		carry, node.Next = sum/10, &ListNode{Val: sum % 10}
		node = node.Next
	}

	if carry > 0 {
		node.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
