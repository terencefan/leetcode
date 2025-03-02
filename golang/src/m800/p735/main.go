package p735

func asteroidCollision(asteroids []int) []int {
	s := make([]int, 0)

	for _, val := range asteroids {
		if val >= 0 {
			s = append(s, val)
		} else {
			for len(s) > 0 && val < 0 {
				if s[len(s)-1] < 0 {
					break
				} else if s[len(s)-1]+val < 0 {
					s = s[:len(s)-1]
				} else if s[len(s)-1]+val == 0 {
					s = s[:len(s)-1]
					val = 0
				} else {
					val = 0
				}
			}
			if val < 0 {
				s = append(s, val)
			}
		}
	}
	return s
}
