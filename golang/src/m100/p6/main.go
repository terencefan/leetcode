package p6

// P   A   H   N
// A P L S I I G
// Y   I   R

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := 2*numRows - 2

	var r = make([]byte, 0)
	for k := 0; k < numRows; k++ {
		if k == 0 {
			for i := 0; i < len(s); i += n {
				r = append(r, s[i])
			}
		} else if k == numRows-1 {
			for i := k; i < len(s); i += n {
				r = append(r, s[i])
			}
		} else {
			for i := 0; i < len(s); i += n {
				if i+k < len(s) {
					r = append(r, s[i+k])
				}
				if i+n-k < len(s) {
					r = append(r, s[i+n-k])
				}
			}
		}
	}
	return string(r)
}
