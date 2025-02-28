package utils

type Queue struct {
	nodes []any
	len   int
	head  int
}

func queue() (q *Queue) {
	q = &Queue{}
	return
}

func (q *Queue) push(node any) {
	if len(q.nodes) == cap(q.nodes) {
		defer func() {
			copy(q.nodes, q.nodes[q.head:])
			q.head = 0
			println("head has been changed!")
		}()
	}
	q.nodes = append(q.nodes, node)
	q.len++
}

func (q *Queue) pop() any {
	if q.len == 0 {
		return nil
	}
	defer func() {
		q.head++
		q.len--
	}()
	return q.nodes[q.head]
}
