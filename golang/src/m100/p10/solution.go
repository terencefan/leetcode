package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if len(s)|len(p) == 0 {
		return true
	}
	match := len(s) > 0 && len(p) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) > 1 && p[1] == '*' {
		return (match && isMatch(s[1:], p)) || isMatch(s, p[2:])
	} else {
		return match && isMatch(s[1:], p[1:])
	}
}

func main() {
	var r bool
	r = isMatch("aab", "c*a*b")
	fmt.Println(r)
	r = isMatch("mississippi", "mis*is*p*.")
	fmt.Println(r)
	r = isMatch("ab", ".*")
	fmt.Println(r)
	r = isMatch("ab", ".*c")
	fmt.Println(r)
	r = isMatch("abcd", "d*")
	fmt.Println(r)
	r = isMatch("aaa", "ab*ac*a")
	fmt.Println(r)
	r = isMatch("aaba", "ab*a*c*a")
	fmt.Println(r)
}
