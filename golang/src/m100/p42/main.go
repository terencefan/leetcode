package p42

func trap(heights []int) int {
	s := []int{}

	var r = 0
	for i, height := range heights {
		for len(s) > 0 && height > heights[s[len(s)-1]] {
			currentTop := s[len(s)-1]
			s = s[:len(s)-1]

			if len(s) == 0 {
				break
			}

			nextTop := s[len(s)-1]
			h := min(height, heights[nextTop]) - heights[currentTop]
			w := i - nextTop - 1
			r += h * w
		}
		s = append(s, i)
	}
	return r
}
