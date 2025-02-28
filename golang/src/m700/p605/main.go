package p605

func canPlace(flowerbed []int, index int) bool {
	if index < 0 || index >= len(flowerbed) {
		return false
	} else if flowerbed[index] == 1 {
		return false
	} else if index > 0 && flowerbed[index-1] == 1 {
		return false
	} else if index < len(flowerbed)-1 && flowerbed[index+1] == 1 {
		return false
	} else {
		return true
	}
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i := range len(flowerbed) {
		if canPlace(flowerbed, i) {
			n--
			flowerbed[i] = 1
			if n == 0 {
				return true
			}
		}
	}
	return false
}
