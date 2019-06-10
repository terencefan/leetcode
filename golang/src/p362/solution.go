package main

type Node struct {
	timestamp, count int
}

type Queue []*Node

func (q *Queue) pop() *Node {
	r := q.head()
	*q = (*q)[1:]
	return r
}

func (q *Queue) push(n *Node) {
	*q = append(*q, n)
}

func (q *Queue) head() *Node {
	return (*q)[0]
}

func (q *Queue) tail() *Node {
	return (*q)[q.Len()-1]
}

func (q *Queue) Len() int {
	return len(*q)
}

type HitCounter struct {
	q     Queue
	count int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		make(Queue, 0),
		0,
	}
}

func (this *HitCounter) adjust(timestamp int) {
	for this.q.Len() > 0 && this.q.head().timestamp <= timestamp-300 {
		node := this.q.pop()
		this.count -= node.count
	}
}

/** Record a hit.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.adjust(timestamp)
	if this.q.Len() > 0 && this.q.tail().timestamp == timestamp {
		this.q.tail().count++
	} else {
		this.q.push(&Node{timestamp, 1})
	}
	this.count++
}

/** Return the number of hits in the past 5 minutes.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	this.adjust(timestamp)
	return this.count
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
