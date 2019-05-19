package main

import (
	"fmt"
)

const ERR = -1.0

type Equations struct {
	m map[string]map[string]float64
}

func (e *Equations) query(a, b string) float64 {
	if e.m[a] == nil {
		return ERR
	}
	if r, ok := e.m[a][b]; ok {
		return r
	} else {
		return ERR
	}
}

func (e *Equations) add(a, b string, v float64) {
	if e.m[a] == nil {
		e.m[a] = make(map[string]float64)
	}
	e.m[a][b] = v

	if e.m[b] == nil {
		e.m[b] = make(map[string]float64)
	}
	e.m[b][a] = 1 / v
}

func (e *Equations) get(a string) map[string]float64 {
	return e.m[a]
}

func query(e *Equations, a, b string, s map[string]bool) float64 {
	r := e.query(a, b)
	if r != ERR {
		return r
	}

	m := e.get(a)
	if m == nil {
		return ERR
	}

	for c, v := range m {
		if s[c] != false {
			continue
		}

		r := e.query(c, b)
		if r != ERR {
			return v * r
		}

		s[a] = true
		r = query(e, c, b, s)
		if r != ERR {
			return v * r
		}
	}
	return ERR
}

func calcEquation(equations [][]string, values []float64, queries [][]string) (r []float64) {
	e := &Equations{
		make(map[string]map[string]float64),
	}

	for i := 0; i < len(equations); i++ {
		e.add(equations[i][0], equations[i][1], values[i])
	}

	r = make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		r[i] = query(e, queries[i][0], queries[i][1], make(map[string]bool))
	}
	return
}

func main() {
	var equations = [][]string{
		{"a", "b"},
		{"b", "c"},
		{"c", "d"},
		{"d", "e"},
	}
	var values = []float64{2.0, 3.0, 4.0, 5.0}
	var queries = [][]string{
		{"a", "e"},
	}
	r := calcEquation(equations, values, queries)
	fmt.Println(r)
}