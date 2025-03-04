package p101

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func isSymmetric(root *TreeNode) bool {
	ql := []*TreeNode{root}
	qr := []*TreeNode{root}

	for len(ql) > 0 && len(qr) > 0 {
		l, r := ql[0], qr[0]
		ql, qr = ql[1:], qr[1:]

		if l == nil && r == nil {
			continue
		} else if l == nil || r == nil {
			return false
		} else if l.Val != r.Val {
			return false
		} else {
			ql = append(ql, l.Left, l.Right)
			qr = append(qr, r.Right, r.Left)
		}
	}
	return len(ql) == len(qr)
}
