package main

func isPowerOfTwo(n int) bool {
	for n > 0 && n&1 == 0 {
		n >>= 1
	}
	return n == 1
}
