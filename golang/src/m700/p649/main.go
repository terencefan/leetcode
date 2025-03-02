package p649

func predictPartyVictory(senate string) string {
	radiant, dire := []int{}, []int{}

	for i, c := range senate {
		if c == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	var ri, di int
	for len(radiant) > 0 && len(dire) > 0 {
		ri, radiant = radiant[0], radiant[1:]
		di, dire = dire[0], dire[1:]

		if ri > di {
			dire = append(dire, di+len(senate))
		} else {
			radiant = append(radiant, ri+len(senate))
		}
	}

	if len(radiant) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}
