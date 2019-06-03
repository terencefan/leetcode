package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	p, q, d := len(num1), len(num2), byte(0)

	b := make([]byte, p)

	for i := p - 1; i >= p - q; i-- {
		b[i] = num2[i - p + q] - '0' + num1[i] + d
		d = 0
		if b[i] > '9' {
			b[i] -= 10
			d = 1
		}
	}

	for i := p - q - 1; i >= 0; i-- {
		b[i] = num1[i] + d
		d = 0
		if b[i] > '9' {
			b[i] -= 10
			d = 1
		}
	}

	if d > 0 {
		return "1" + string(b)
	} else {
		return string(b)
	}
}

func main() {
	fmt.Println(addStrings("9", "99"))
}