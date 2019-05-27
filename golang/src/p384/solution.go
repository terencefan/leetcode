package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	arr []int
	origin []int
}


func Constructor(nums []int) Solution {
	origin := make([]int, len(nums))
	copy(origin, nums)
	return Solution{
		arr: nums,
		origin: origin,
	}
}


/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	copy(s.arr, s.origin)
	return s.arr
}


/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	for i := 0; i <len(s.arr); i++ {
		next := rand.Intn(len(s.arr) - i) + i
		s.arr[i], s.arr[next] = s.arr[next], s.arr[i]
	}
	return s.arr
}

func main() {
	s := Constructor([]int{1,2,3})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
}
