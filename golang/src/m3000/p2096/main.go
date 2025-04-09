package p2096

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

type Edge struct {
	val       int
	direction rune
}

func rec(node *TreeNode, m map[int][]Edge) {
	if node.Left != nil {
		m[node.Val] = append(m[node.Val], Edge{node.Left.Val, 'L'})
		m[node.Left.Val] = append(m[node.Left.Val], Edge{node.Val, 'U'})
		rec(node.Left, m)
	}
	if node.Right != nil {
		m[node.Val] = append(m[node.Val], Edge{node.Right.Val, 'R'})
		m[node.Right.Val] = append(m[node.Right.Val], Edge{node.Val, 'U'})
		rec(node.Right, m)
	}

}

func dfs(m map[int][]Edge, visited map[int]bool, bytes *[]rune, val, dest int) bool {
	if val == dest {
		return true
	}
	visited[val] = true

	for _, next := range m[val] {
		*bytes = append(*bytes, next.direction)
		if !visited[next.val] && dfs(m, visited, bytes, next.val, dest) {
			return true
		}
		*bytes = (*bytes)[:len(*bytes)-1]
	}
	visited[val] = false
	return false
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var m = map[int][]Edge{}
	rec(root, m)

	var bytes = make([]rune, 0)
	var visited = map[int]bool{}
	dfs(m, visited, &bytes, startValue, destValue)
	return string(bytes)
}
