package p841

func visit(rooms [][]int, visited []bool, room int) int {
	if visited[room] {
		return 0
	}
	visited[room] = true
	r := 1

	for _, key := range rooms[room] {
		r += visit(rooms, visited, key)
	}
	return r
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)

	if n < 2 {
		return true
	}

	for i := range len(rooms) {
		visited := make([]bool, n)
		if n == visit(rooms, visited, i) {
			return true
		}
	}
	return false
}
