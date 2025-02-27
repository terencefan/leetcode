package main

import (
	"fmt"
	"strconv"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func str(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	} else {
		return fmt.Sprintf("%d->%d", a, b)
	}
}

func findMissingRanges(nums []int, lower int, upper int) (r []string) {
	for _, num := range nums {
		if num > upper {
			break
		}
		if num < lower {
			continue
		} else if num == lower {
		} else {
			r = append(r, str(lower, min(num-1, upper)))
		}
		lower = num + 1
	}
	if lower <= upper {
		r = append(r, str(lower, upper))
	}
	return
}

func main() {
	r := findMissingRanges([]int{0, 1, 3, 50, 75}, 6, 72)
	fmt.Println(r)
}
