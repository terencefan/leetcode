package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var r = make([]bool, 0)

	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	for _, candy := range candies {
		r = append(r, candy+extraCandies >= maxCandies)
	}
	return r
}
