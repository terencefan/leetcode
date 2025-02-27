package main

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.requests) > 0 && this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}
	this.requests = append(this.requests, t)
	return len(this.requests)
}
