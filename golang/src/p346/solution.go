package main

type MovingAverage struct {
	arr               []int
	pos, count, total int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		make([]int, size),
		0, 0, 0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.count < len(this.arr) {
		this.count++
	}
	pos := this.pos % len(this.arr)
	this.pos++
	this.total += val - this.arr[pos]
	this.arr[pos] = val
	return float64(this.total) / float64(this.count)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
