package main

func IsPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}
	var t = num / 2
	for t*t > num {
		t = (t + num/t) / 2
	}
	return t*t == num
}
