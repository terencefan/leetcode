package p155

type MinStack struct {
	all, min []int
	count    int
}

func Constructor() MinStack {
	return MinStack{
		all: make([]int, 0),
		min: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.all = append(this.all, val)
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	val := this.all[len(this.all)-1]
	this.all = this.all[:len(this.all)-1]
	if len(this.min) > 0 && val == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.all[len(this.all)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
