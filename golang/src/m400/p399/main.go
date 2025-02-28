package p399

type Denominator struct {
	name string
	val  float64
}
type FactorMap map[string][]Denominator

func dfs(m FactorMap, a, b string, visited map[string]bool) (float64, bool) {
	if a == b {
		return 1.0, true
	}

	for _, next := range m[a] {
		if visited[next.name] {
			continue
		}
		visited[next.name] = true

		if val, ok := dfs(m, next.name, b, visited); ok {
			return val * next.val, true
		}
	}
	return -1, false
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var m = make(FactorMap)

	for i, equation := range equations {
		a, b := equation[0], equation[1]
		m[a] = append(m[a], Denominator{name: b, val: values[i]})
		m[b] = append(m[b], Denominator{name: a, val: 1.0 / values[i]})
	}

	var res = make([]float64, len(queries))
	for i, query := range queries {
		var visited = make(map[string]bool)

		a, b := query[0], query[1]

		if _, ok := m[a]; !ok {
			res[i] = -1.0
		} else {
			res[i], _ = dfs(m, a, b, visited)
		}
	}
	return res
}
