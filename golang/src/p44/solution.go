package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	r := []byte{p[0]}
	for i := 1; i < len(p); i++ {
		if r[len(r)-1] == '*' {
			if p[i] == '*' {
				continue
			}
		}
		r = append(r, p[i])
	}

	i, j, ti, tj := 0, 0, -1, -1
	for i < len(s) {
		if j < len(r) && (s[i] == r[j] || r[j] == '?') {
			i++
			j++
			continue
		} else if j < len(r) && r[j] == '*' {
			ti = i
			tj = j
			j++
		} else if ti < 0 {
			return false
		} else {
			i, j = ti+1, tj+1
			ti = i
		}
	}
	for k := j; k < len(r); k++ {
		if r[k] != '*' {
			return false
		}
	}
	return true
}

func main() {
	r := isMatch(
		"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
		"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb",
	)
	fmt.Println(r)

	fmt.Println(isMatch("mississippi", "m??*ss*?i*pi"))

	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("aaaa", "***a"))
	fmt.Println(isMatch("", ""))
}
