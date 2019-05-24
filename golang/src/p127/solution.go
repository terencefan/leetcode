package main

import (
	"fmt"
)

func hasRoute(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	c := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if c == 1 {
				return false
			}
			c++
		}
	}
	return true
}

type graph [][]int

func (g *graph) add(x, y int) {
	if (*g)[x] == nil {
		(*g)[x] = []int{y}
	} else {
		(*g)[x] = append((*g)[x], y)
	}
}

func (g *graph) routes(x int) []int {
	return (*g)[x]
}

func (g *graph) len() int {
	return len(*g)
}

func copy(m map[int]bool, x int) (n map[int]bool) {
	n = make(map[int]bool, len(m) + 1)
	for k, v := range m {
		n[k] = v
	}
	n[x] = true
	return
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var m = make(map[string]int)
	var g = make(graph, len(wordList) + 1)
	var endIndex = 0

	for i, word := range wordList {
		m[word] = i + 1
		if hasRoute(beginWord, word) {
			g.add(0, i + 1)
		}
		if word == endWord {
			endIndex = i + 1
		}
	}
	if endIndex == 0 {
		return 0
	}

	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			w1, w2 := wordList[i], wordList[j]
			if hasRoute(w1, w2) {
				g.add(m[w1], m[w2])
				g.add(m[w2], m[w1])
			}
		}
	}

	q := []int{0}
	visited := map[int]bool{0: true}

	r := 1
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[i]
			if node == endIndex {
				return r
			}
			for _, next := range g.routes(node) {
				if visited[next] == true {
					continue
				}
				visited[next] = true
				q = append(q, next)
			}
		}
		q = q[l:]
		r++
	}
	return 0
}

func main() {
	r := ladderLength(
		"hit",
		"cog",
		[]string{"hot","dot","dog","lot","log","cog"},
	)
	fmt.Println(r)
}
