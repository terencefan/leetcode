package main

type P int

func (p P) getX() int {
	return int(p&0xFF00) >> 8
}

func (p P) getY() int {
	return int(p & 0xFF)
}

func New(x, y int) P {
	return P(x<<8 + y)
}

var moves = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func isValid(x, y, height, width int) bool {
	return !(x < 0 || y < 0 || x == height || y == width)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}
	height, width := len(image), len(image[0])

	if !isValid(sr, sc, height, width) {
		return image
	}
	oldColor := image[sr][sc]

	if newColor == oldColor {
		return image
	}

	p, q := P(0), []P{New(sr, sc)}
	for len(q) > 0 {
		p, q = q[0], q[1:]
		x, y := p.getX(), p.getY()

		if image[x][y] != oldColor {
			continue
		}
		image[x][y] = newColor

		for _, move := range moves {
			nx, ny := x+move[0], y+move[1]
			if nx < 0 || ny < 0 || nx == height || ny == width {
				continue
			}
			q = append(q, New(nx, ny))
		}
	}

	return image
}
