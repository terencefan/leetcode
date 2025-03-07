package p117

import "terencefan.com/leetcode/src/utils"

type Node = utils.TreeNodeNext

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := []*Node{root}

	var prev, cur *Node
	for len(q) > 0 {
		l := len(q)

		prev = nil
		for i := range l {
			cur = q[i]
			if prev != nil {
				prev.Next = cur
			}
			prev = cur

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if q[i].Right != nil {
				q = append(q, cur.Right)
			}
		}

		q = q[l:]
	}
	return root
}
