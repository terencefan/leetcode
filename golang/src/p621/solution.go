package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func leastInterval1(tasks []byte, n int) int {
	var taskArr = make([]int, 26)
	for _, task := range tasks {
		taskArr[int(task-'A')]++
	}
	sort.Ints(taskArr)

	round := func() (r int) {
		m := n + 1
		for i := 25; i >= 0 && m > 0 && taskArr[i] > 0; {
			j := i
			for ; j >= 0; j-- {
				if taskArr[j] < taskArr[i] {
					for k := 0; k < i-j && k < m; k++ {
						taskArr[j+k+1]--
						r++
					}
					m -= i - j
					break
				}
			}
			i = j
		}
		return
	}

	r, last := 0, 0
	for taskArr[len(taskArr)-1] > 0 {
		r++
		last = round()
	}
	return (r-1)*(n+1) + last
}

func leastInterval(tasks []byte, n int) int {
	var taskArr = make([]int, 26)
	for _, task := range tasks {
		taskArr[int(task-'A')]++
	}
	sort.Ints(taskArr)

	max := taskArr[25]
	idle := (max - 1) * n

	for i := 24; i >= 0 && taskArr[i] > 0; i-- {
		if taskArr[i] == max {
			idle++
		}
		idle -= taskArr[i]
	}

	if idle < 0 {
		return len(tasks)
	} else {
		return idle + len(tasks)
	}
}

func main() {
	cases := make([]byte, 0)

	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		b := byte(int(math.Sqrt(float64(ran.Intn(26*26)))) + int('A'))
		cases = append(cases, b)
	}

	r := leastInterval1(cases, 43)
	fmt.Println(r)
}
