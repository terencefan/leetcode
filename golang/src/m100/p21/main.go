package p21

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var cur = dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next, cur = list1, list1
			list1 = list1.Next
		} else {
			cur.Next, cur = list2, list2
			list2 = list2.Next
		}
	}

	for list1 != nil {
		cur.Next, cur = list1, list1
		list1 = list1.Next
	}

	for list2 != nil {
		cur.Next, cur = list2, list2
		list2 = list2.Next
	}
	return dummy.Next
}
