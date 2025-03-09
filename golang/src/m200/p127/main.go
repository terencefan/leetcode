package p127

func canTransform(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := 0

	for i := range a {
		if a[i] != b[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	m := make([][]int, n)
	for i := range n {
		for j := i + 1; j < n; j++ {
			if canTransform(wordList[i], wordList[j]) {
				m[i] = append(m[i], j)
				m[j] = append(m[j], i)
			}
		}
	}

	var q = make([]int, 0)
	var visited = make([]bool, n)
	var endIdx = -1

	for idx, word := range wordList {
		if word == endWord {
			endIdx = idx
		}
		if canTransform(beginWord, word) {
			q = append(q, idx)
			visited[idx] = true
		}
	}
	if endIdx < 0 {
		return 0
	} else if canTransform(beginWord, endWord) {
		return 2
	}

	r := 2
	for len(q) > 0 {
		l := len(q)

		for i := range l {
			idx := q[i]
			for _, nextIdx := range m[idx] {
				if nextIdx == endIdx {
					return r + 1
				} else if !visited[nextIdx] {
					q = append(q, nextIdx)
					visited[nextIdx] = true
				}
			}
		}
		q = q[l:]
		r++
	}
	return 0
}
