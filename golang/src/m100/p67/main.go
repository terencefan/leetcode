package p67

func getByte(a string, i int) byte {
	if i < len(a) {
		return a[len(a)-1-i]
	} else {
		return '0'
	}
}

func addBinary(a string, b string) string {
	var res = make([]byte, max(len(a), len(b))+1)
	var add = false

	for p := range res {
		pos := len(res) - p - 1
		x, y := getByte(a, p), getByte(b, p)
		if x == '1' && y == '1' {
			if add {
				res[pos] = '1'
			} else {
				res[pos] = '0'
				add = true
			}
		} else if x == '0' && y == '0' {
			if add {
				res[pos] = '1'
			} else {
				res[pos] = '0'
			}
			add = false
		} else {
			if add {
				res[pos] = '0'
			} else {
				res[pos] = '1'
			}
		}
	}
	if len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return string(res)
}
