package p12

var intMap = map[int]byte{
	1:    'I',
	5:    'V',
	10:   'X',
	50:   'L',
	100:  'C',
	500:  'D',
	1000: 'M',
}

func digitToBytes(d int, one, five, ten byte) []byte {
	if d > 3 && five == 0 {
		panic("invalid digit")
	} else if d > 8 && ten == 0 {
		panic("invalid digit")
	} else if d < 0 || d > 9 {
		panic("invalid digit")
	}

	r := make([]byte, 0)
	if d%5 == 4 {
		r = append(r, one)
	}
	if d > 8 {
		r = append(r, ten)
	} else if d > 3 {
		r = append(r, five)
	}
	if d%5 != 4 {
		for range d % 5 {
			r = append(r, one)
		}
	}
	return r
}

func intToRoman(num int) string {
	var r = make([]byte, 0)
	for num > 0 {
		if num >= 1000 {
			r = append(r, digitToBytes(num/1000, 'M', 0, 0)...)
			num -= num / 1000 * 1000
		} else if num >= 100 {
			r = append(r, digitToBytes(num/100, 'C', 'D', 'M')...)
			num -= num / 100 * 100
		} else if num >= 10 {
			r = append(r, digitToBytes(num/100, 'X', 'L', 'C')...)
			num -= num / 10 * 10
		} else if num >= 1 {
			r = append(r, digitToBytes(num/100, 'I', 'V', 'X')...)
			num = 0
		}
	}
	return string(r)
}
