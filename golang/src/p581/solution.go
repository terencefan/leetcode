package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) peek() int {
	return (*s)[len(*s) - 1]
}

func (s *Stack) push(x int) {
	*s = append(*s, x)
}

func (s *Stack) pop() int {
	l := s.len()
	defer func() {
		*s = (*s)[:l - 1]
	}()
	return (*s)[l - 1]
}

func (s *Stack) len() int {
	return len(*s)
}

func findLeft(nums *[]int) int {
	l := len(*nums) - 1
	i := 1
	for ; i < len(*nums); i++ {
		if (*nums)[i] < (*nums)[i - 1] {
			l = i - 1
			break
		}
	}

	for ;i < len(*nums) && l >= 0; i++ {
		for l >= 0 && (*nums)[i] < (*nums)[l] {
			l--
		}
	}
	return l
}

func findRight(nums *[]int) int {
	r := 0
	i := len(*nums) - 2

	for ; i >= 0; i-- {
		if (*nums)[i] > (*nums)[i + 1] {
			r = i + 1
			break
		}
	}

	for ; i >= 0 && r < len(*nums); i-- {
		for r < len(*nums) && (*nums)[i] > (*nums)[r] {
			r++
		}
	}
	return r
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	left := findLeft(&nums)
	right := findRight(&nums)
	if right < left {
		return 0
	}
	return right - left - 1
}

func main() {
	r := findUnsortedSubarray([]int{5, 2, 3, 6})
	fmt.Println(r)
}