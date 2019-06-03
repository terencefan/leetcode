package main

func isPowerOfFour(num int) bool {
	for num > 0 && num & 3 == 0 {
		num >>= 2
	}
	return num == 1
}
