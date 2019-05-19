package main

import (
	"fmt"
)

const LP = '('
const RP = ')'

type Queue struct {
	l, r int
	cap  int
	arr  []*Node
}

func (q *Queue) push(node *Node) {
	if q.r - q.l == q.cap {
		oldArr := q.arr
		q.arr = make([]*Node, q.cap* 2)

		k := 0
		for i := q.l; i < q.r; i++ {
			q.arr[k] = oldArr[i % q.cap]
			k++
		}
		q.l = 0
		q.r = q.cap
		q.cap *= 2
	}
	q.arr[q.r % q.cap] = node
	q.r++
}

func (q *Queue) pop() *Node {
	defer func() {q.l++}()
	return q.arr[q.l % q.cap]
}

func (q *Queue) empty() bool {
	return q.l == q.r
}

func NewQueue() *Queue {
	return &Queue{
		0, 0, 16,
		make([]*Node, 16),
	}
}

type Node struct {
	s string
	l, r int
	l2r bool
}

func bfs(q *Queue, r *[]string) {
	if q.empty() {
		return
	}

	node := q.pop()
	var count int

	if node.l2r {
		for i := node.l; i < len(node.s); i++ {
			if node.s[i] == LP {
				count++
			} else if node.s[i] == RP {
				count--
			}

			if count < 0 {
				for j := node.r; j <= i; j++ {
					if node.s[j] == RP && (j == node.r || node.s[j-1] != RP) {
						q.push(&Node{node.s[:j] + node.s[j+1:], i, j, true})
						bfs(q, r)
					}
				}
				return
			}
		}
		q.push(&Node{node.s, len(node.s) - 1, len(node.s) -1, false})
		bfs(q, r)
	} else {
		for i := node.l; i >= 0; i-- {
			if node.s[i] == LP {
				count--
			} else if node.s[i] == RP {
				count++
			}

			if count < 0 {
				for j := node.r; j >= i; j-- {
					if node.s[j] == LP && (j == node.r || node.s[j+1] != LP) {
						q.push(&Node{node.s[:j] + node.s[j+1:], i - 1, j - 1, false})
						bfs(q, r)
					}
				}
				return
			}

		}
		*r = append(*r, node.s)
	}
	return
}

func removeInvalidParentheses(s string) []string {
	var r []string

	q := NewQueue()
	q.push(&Node{s, 0, 0, true})

	bfs(q, &r)
	return r
}

func main() {
	r := removeInvalidParentheses("(()")
	fmt.Println(r)
}
