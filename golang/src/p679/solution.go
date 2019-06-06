package main

import (
	"fmt"
)

const (
	ADD = iota + 10
	SUB
	MUL
	DIV
)

func abs(a float64) float64 {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func calculate(items []int) float64 {
	var s []float64
	for _, item := range items {
		if item > 9 {
			if len(s) < 2 {
				return -1
			}
			a, b := s[len(s)-1], s[len(s)-2]
			s = s[:len(s)-1]
			switch item {
			case ADD:
				s[len(s)-1] = a + b
			case SUB:
				s[len(s)-1] = a - b
			case MUL:
				s[len(s)-1] = a * b
			case DIV:
				if b == 0 {
					return -1
				}
				s[len(s)-1] = a / float64(b)
			}
		} else {
			s = append(s, float64(item))
		}
	}
	return s[0]
}

func arrangement(nums, items []int, index, count int) bool {
	if index == 7 {
		if abs(calculate(items)-24) < 1e-9 {
			return true
		}
	}

	if count > 1 {
		for i := ADD; i <= DIV; i++ {
			items[index] = i
			if arrangement(nums, items, index+1, count-1) {
				return true
			}
		}
	}

	var t int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			t, nums[i] = nums[i], 0
			items[index] = t
			if arrangement(nums, items, index+1, count+1) {
				return true
			}
			nums[i] = t
		}
	}
	return false
}

func judgePoint24(nums []int) bool {
	items := make([]int, 7)
	if arrangement(nums, items, 0, 0) {
		return true
	}
	return false
}

func main() {
	r := judgePoint24([]int{4, 1, 8, 7})
	fmt.Println(r)
}
