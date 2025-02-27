package main

type Bed []int

func (b Bed) canPlace(i int) bool {
	if len(b) == 0 {
		return false
	} else if i >= len(b) {
		return false
	}

	if len(b) == 1 {
		return i == 0 && b[0] == 0
	} else if i == 0 {
		return b[i]|b[i+1] == 0
	} else if i == len(b)-1 {
		return b[i]|b[i-1] == 0
	} else {
		return b[i]|b[i-1]|b[i+1] == 0
	}
}

func canPlaceFlowers(flowerbed Bed, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed.canPlace(i) {
			flowerbed[i] = 1
			n--
		}
	}
	return n == 0
}
