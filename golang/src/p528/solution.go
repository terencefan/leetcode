package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	arr []int
	max int
}

func Constructor(w []int) Solution {
	arr := make([]int, 0)

	c := 0
	for _, num := range w {
		c += num
		arr = append(arr, c)
	}

	return Solution{
		arr: arr,
		max: c,
	}
}

func (this *Solution) PickIndex() int {
	current := rand.Intn(this.max)

	lo, hi := 0, len(this.arr)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if current >= this.arr[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	s := Constructor([]int{1, 2, 3})
	for i := 0; i < 10; i++ {
		fmt.Println(s.PickIndex())
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
