package p49

const MOD = 1e9 + 7

var primes = map[rune]int{
	'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19, 'i': 23, 'j': 29,
	'k': 31, 'l': 37, 'm': 41, 'n': 43, 'o': 47, 'p': 53, 'q': 59, 'r': 61, 's': 67, 't': 71,
	'u': 73, 'v': 79, 'w': 83, 'x': 89, 'y': 97, 'z': 101,
}

func hash(s string) int {
	r := 1
	for _, v := range s {
		r *= primes[v]
		r %= MOD
	}
	return r
}

func groupAnagrams(strs []string) [][]string {
	var m = make(map[int][]string)

	for _, str := range strs {
		code := hash(str)
		m[code] = append(m[code], str)
	}

	var r = make([][]string, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
