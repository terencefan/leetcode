package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	minus := false
	if numerator == 0 {
		return "0"
	} else if numerator < 0 {
		minus = !minus
		numerator = -numerator
	}
	if denominator < 0 {
		minus = !minus
		denominator = -denominator
	}

	var intp = make([]byte, 0)
	if minus {
		intp = append(intp, '-')
	}
	if numerator < denominator {
		intp = append(intp, '0')
	} else {
		intp = append(intp, strconv.Itoa(numerator/denominator)...)
	}

	numerator = numerator % denominator
	if numerator == 0 {
		return string(intp)
	}

	var intf = make([]byte, 0)
	var m = make(map[int]int)

	pos := 0
	for numerator > 0 {
		numerator *= 10
		v := numerator / denominator
		numerator = numerator % denominator

		r := m[numerator*10+v]
		if r > 0 {
			result := string(intp) + "." + string(intf[:r-1]) + "(" + string(intf[r-1:]) + ")"
			return result
		}
		intf = append(intf, byte(v+'0'))
		pos++
		m[numerator*10+v] = pos
	}
	return string(intp) + "." + string(intf)
}

func main() {
	r := fractionToDecimal(500, 10)
	fmt.Println(r)
}
