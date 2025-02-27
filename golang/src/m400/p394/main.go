package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func decode(s string, index int) (int, []byte) {
	var bytes = make([]byte, 0)
	var repeat = 0

	for i := index; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			repeat = repeat*10 + int(s[i]-'0')
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			bytes = append(bytes, s[i])
		case '[':
			var sub []byte
			i, sub = decode(s, i+1)
			for k := 0; k < max(repeat, 1); k++ {
				bytes = append(bytes, sub...)
			}
			repeat = 0
		case ']':
			return i, bytes
		}
	}
	return len(s), bytes
}

func decodeString(s string) string {
	_, bytes := decode(s, 0)
	return string(bytes)
}

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}
