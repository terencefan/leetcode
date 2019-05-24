package main

import (
	"fmt"
)

func fizzBuzz(n int) (r []string) {
	for i := 0; i < n; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				r = append(r, "FizzBuzz")
			} else {
				r = append(r, "Fizz")
			}
		} else if i % 5 == 0 {
			r = append(r, "Buzz")
		} else {
			r = append(r, fmt.Sprintf("%d", i))
		}
	}
	return r
}