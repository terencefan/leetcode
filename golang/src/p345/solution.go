package main

import (
	"fmt"
)

func isVowel(b byte) bool {
	switch b {
	case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	b := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for ; i < j && !isVowel(b[i]); i++ {
		}
		for ; i < j && !isVowel(b[j]); j-- {
		}
		if i != j {
			b[i], b[j] = b[j], b[i]
		}
		i++
		j--
	}
	return string(b)
}

func main() {
	fmt.Println(reverseVowels("hello"))
}
