package p13

var byteMap = map[byte]int{
	' ': 0,
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	pb := s[0]
	r, c := 0, byteMap[pb]

	for i := 1; i < len(s); i++ {
		b := s[i]
		if pb == b {
			c += byteMap[b]
		} else if byteMap[b] > byteMap[pb] {
			r += byteMap[b] - c
			c = 0
		} else {
			r += c
			c = byteMap[b]
		}
		pb = b
	}
	return r + c
}
