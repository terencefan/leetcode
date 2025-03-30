package p1010

func numPairsDivisibleBy60(time []int) int {

	mods := make(map[int]int)

	for _, t := range time {
		mods[t%60]++
	}

	var r = 0
	for i := 1; i < 30; i++ {
		r += mods[i] * mods[60-i]
	}
	r += mods[0] * (mods[0] - 1) / 2
	r += mods[30] * (mods[30] - 1) / 2
	return r
}
