package p427

import "terencefan.com/leetcode/src/utils"

type Node = utils.QuadTreeNode

func build(grid [][]int, x, y, size int) *Node {
	if size == 1 {
		return &Node{Val: grid[x][y] == 1, IsLeaf: true}
	}

	tl := build(grid, x, y, size/2)
	tr := build(grid, x, y+size/2, size/2)
	bl := build(grid, x+size/2, y, size/2)
	br := build(grid, x+size/2, y+size/2, size/2)

	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val {
		return &Node{Val: tl.Val, IsLeaf: true}
	} else {
		return &Node{TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
	}
}

func construct(grid [][]int) *Node {
	return build(grid, 0, 0, len(grid))
}
