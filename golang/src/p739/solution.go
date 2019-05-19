package main

import (
	"fmt"
)

func pack(pos, t int) int {
	return t << 16 + pos
}

func unpack(v int) (pos, t int) {
	pos = v & 0x0000ffff
	t = (v & 0xffff0000) >> 16
	return
}

func dailyTemperatures(T []int) []int {
	length := len(T)
	r := make([]int, length)

	stack := make([]int, 0, 16)
	for i := length - 1; i >=0 ; i-- {
		for len(stack) > 0 {
			pos, t := unpack(stack[len(stack) - 1])
			if t < T[i] {
				stack = stack[:len(stack) - 1]
			} else {
				r[i] = pos - i
				break
			}
		}
		stack = append(stack, pack(i, T[i]))
	}

	return r
}

func main() {
	r := dailyTemperatures([]int{73,74,75,71,69,72,76,73})
	fmt.Println(r)
}