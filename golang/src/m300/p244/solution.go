package main

import "fmt"

type WordDistance struct {
	words map[string][]int
}

func Constructor(words []string) WordDistance {
	wordMap := make(map[string][]int)
	for index, word := range words {
		if positions, ok := wordMap[word]; ok {
			wordMap[word] = append(positions, index)
		} else {
			wordMap[word] = []int{index}
		}
	}
	return WordDistance{
		words: wordMap,
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	// TODO add a cache

	var (
		p1, p2 []int
		ok     bool
	)

	if p1, ok = this.words[word1]; !ok {
		return 0
	}

	if p2, ok = this.words[word2]; !ok {
		return 0
	}

	if len(p1) == 0 || len(p2) == 0 {
		return 0
	}

	i, j, minDistance := 0, 0, abs(p1[0]-p2[0])
	for i < len(p1) && j < len(p2) {
		minDistance = min(minDistance, abs(p1[i]-p2[j]))
		if p1[i] < p2[j] {
			i++
		} else {
			j++
		}
	}
	return minDistance
}

func main() {
	d := Constructor([]string{
		"apple",
		"haha",
		"banana",
		"apple",
		"apple",
	})
	fmt.Println(d.Shortest("apple", "banana"))
}
