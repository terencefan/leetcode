package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func diff(l, r string, g *G) {
	for i := 0; i < min(len(l), len(r)); i++ {
		if l[i] != r[i] {
			g.add(l[i]-'a', r[i]-'a')
			break
		}
	}
	return
}

type G struct {
	v map[byte][]byte
}

func (g *G) add(a, b byte) {
	if g.v[a] == nil {
		g.v[a] = []byte{b}
	} else {
		g.v[a] = append(g.v[a], b)
	}
}

func topologySort(g *G, chars []int) []byte {

	count := 0
	for i := 0; i < 26; i++ {
		if chars[i] < 0 {
			chars[i] = 0
			count++
		} else {
			chars[i] = -1
		}
	}

	for _, arr := range g.v {
		for _, r := range arr {
			chars[r]++
		}
	}

	cur := byte(26)
	res := make([]byte, 0)

	for {
		if cur > 25 {
			for i := 0; i < 26; i++ {
				if chars[i] == 0 {
					cur = byte(i)
				}
			}
		}
		if cur > 25 {
			break
		}
		res = append(res, byte(cur+'a'))
		chars[cur] = -1

		nexts := g.v[cur]
		cur = 26

		for _, r := range nexts {
			chars[r]--
			if chars[r] == 0 {
				cur = r
			}
		}
	}

	if count == len(res) {
		return res
	} else {
		return []byte{}
	}
}

func alienOrder(words []string) string {
	var g = &G{make(map[byte][]byte)}

	for i := 0; i < len(words)-1; i++ {
		diff(words[i], words[i+1], g)
	}

	var chars = make([]int, 26)
	for _, word := range words {
		for _, c := range word {
			chars[c-'a'] = -1
		}
	}

	b := topologySort(g, chars)
	return string(b)
}

func main() {
	r := alienOrder([]string{
		"wrt",
		"wrf",
	})
	fmt.Println(r)
}
