package main

import (
	"fmt"
	"strconv"
)

func toString(s, t int) string {
	if s == t {
		return strconv.Itoa(s)
	} else {
		return fmt.Sprintf("%d->%d", s, t)
	}
}

func summaryRanges(nums []int) (r []string) {
	if len(nums) == 0 {
		return []string{}
	}

	s, t := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num == t + 1 {
			t = num
		} else {
			r = append(r, toString(s, t))
			s, t = num, num
		}
	}
	r = append(r, toString(s, t))
	return
}

func main() {
	r := summaryRanges([]int{0,1,2,4,5,7})
	fmt.Println(r)
}