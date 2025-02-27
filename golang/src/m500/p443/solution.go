package main

import "fmt"

func flush(chars []byte, index int, count int, last byte) int {
	var n = 1
	for n <= count {
		n *= 10
	}

	chars[index] = last
	index++

	if count == 1 {
		return index
	}

	for n /= 10; n > 0; n /= 10 {
		chars[index] = byte(count/n) + '0'
		count %= n
		index++
	}
	return index
}

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}

	var lp, last, count = 0, chars[0], 1

	for rp := 1; rp < len(chars); rp++ {
		if chars[rp] != last {
			lp = flush(chars, lp, count, last)
			last = chars[rp]
			count = 1
		} else {
			count++
		}
	}
	return flush(chars, lp, count, last)
}

func main() {
	var r = []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(r))
	fmt.Println(string(r))
}
