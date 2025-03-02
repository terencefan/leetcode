package p739

type IndexTemp struct {
	index       int
	temperature int
}

const INTMAX = 1 << 31

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	s := []IndexTemp{{n, INTMAX}}

	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]

		for s[len(s)-1].temperature <= t {
			s = s[:len(s)-1]
		}

		peak := s[len(s)-1]
		if peak.temperature == INTMAX {
			temperatures[i] = 0
		} else {
			temperatures[i] = peak.index - i
		}
		s = append(s, IndexTemp{i, t})
	}
	return temperatures
}
