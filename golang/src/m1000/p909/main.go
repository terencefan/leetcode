package p909

const INTMAX = 1<<31 - 1

func normalize2(idx, n int) (int, int) {
	dx := n - idx/n - 1
	dy := idx % n

	if dx%2^n%2 == 0 {
		return dx, n - dy - 1
	} else {
		return dx, dy
	}
}

func normalize(n int, l int) (int, int) {
	x := n / l
	y := n % l
	if x%2 == 1 {
		y = l - 1 - y
	}
	return l - 1 - x, y
}

func snakesAndLadders(board [][]int) int {
	var n = len(board)
	var n2 = n * n
	var ladders = make(map[int]int)

	for i := range n * n {
		x, y := normalize(i, n)

		if board[x][y] != -1 {
			ladders[i+1] = board[x][y]
		}
	}

	var q = [][]int{{1, 0}}
	var visited = make([]bool, n2)

	for len(q) > 0 {
		node := q[0]
		idx, steps := node[0], node[1]
		q = q[1:]

		if idx == n2 {
			return steps
		} else if visited[idx] {
			continue
		}
		visited[idx] = true

		for k := idx + 1; k <= idx+6 && k <= n2; k++ {
			if to, ok := ladders[k]; ok {
				q = append(q, []int{to, steps + 1})
			} else {
				q = append(q, []int{k, steps + 1})
			}
		}
	}
	return -1
}
