package p433

func canMutate(from, to string) bool {
	if len(from) != len(to) {
		return false
	}

	var diff = 0
	for i := range from {
		if from[i] != to[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

type IndexCount struct {
	index, count int
}

func minMutation(startGene string, endGene string, bank []string) int {
	var n = len(bank)
	var endIdx = -1

	var mutate = make([][]int, n)
	for i := range bank {
		mutate[i] = make([]int, 0)
		if bank[i] == endGene {
			endIdx = i
		}
	}
	if endIdx < 0 {
		return -1
	}

	for i := range bank {
		for j := range bank {
			if i != j && canMutate(bank[i], bank[j]) {
				mutate[i] = append(mutate[i], j)
				mutate[j] = append(mutate[j], i)
			}
		}
	}

	var visited = make([]bool, n)

	var q = make([]IndexCount, 0)
	for i, gene := range bank {
		if canMutate(startGene, gene) {
			q = append(q, IndexCount{i, 1})
			visited[i] = true
		}
	}

	for len(q) > 0 {
		l := len(q)
		for i := range l {
			idx := q[i].index
			if idx == endIdx {
				return q[i].count
			}
			for _, nextIdx := range mutate[idx] {
				if !visited[nextIdx] {
					q = append(q, IndexCount{nextIdx, q[i].count + 1})
				}
			}
		}
		q = q[l:]
	}
	return -1
}
