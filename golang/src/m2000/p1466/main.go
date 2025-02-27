package main

import "fmt"

type Road struct {
	dest  int
	value int
}

func dfs(src int, roads [][]Road, visited []bool) int {
	visited[src] = true
	var r = 0
	for _, road := range roads[src] {
		if !visited[road.dest] {
			r += dfs(road.dest, roads, visited) + road.value
		}
	}
	return r
}

func minReorder(n int, connections [][]int) int {
	if n <= 1 {
		return 0
	}

	var roads = make([][]Road, n)
	var visited = make([]bool, n)

	for _, connection := range connections {
		var start, end = connection[0], connection[1]
		roads[start] = append(roads[start], Road{end, 1})
		roads[end] = append(roads[end], Road{start, 0})
	}

	return dfs(0, roads, visited)
}

func case1() {
	var connections = [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	var r = minReorder(6, connections)
	fmt.Println(r)
}

func case2() {
	var connections = [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}
	var r = minReorder(5, connections)
	fmt.Println(r)
}

func main() {
	case1()
	case2()
}
