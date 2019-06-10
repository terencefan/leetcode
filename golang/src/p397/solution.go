package main

func integerReplacement(n int) int {
	q := []int{n}

	step := 0
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			v := q[i]
			if v == 1 {
				return step
			}

			if v%2 == 0 {
				q = append(q, v/2)
			} else {
				q = append(q, v-1)
				q = append(q, v+1)
			}
		}
		q = q[length:]
		step++
	}
	return -1
}
