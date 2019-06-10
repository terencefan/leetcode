package main

func fib(N int) int {
	if N <= 0 {
		return 0
	} else if N == 1 {
		return 1
	}

	f1, f2 := 0, 1
	for N > 1 {
		f1, f2 = f2, f1+f2
		N--
	}
	return f2
}
