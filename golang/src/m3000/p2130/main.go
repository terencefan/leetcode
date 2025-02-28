package p2130

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pairSum(head *ListNode) int {

	lp, rp := head, head

	var prev *ListNode = nil

	for rp != nil {
		next := lp.Next
		rp = rp.Next.Next

		lp.Next = prev
		prev = lp
		lp = next
	}

	var r = 0
	for lp != nil && prev != nil {
		sum := lp.Val + prev.Val
		r = max(r, sum)
		lp = lp.Next
		prev = prev.Next
	}
	return r
}
