package main

func countSegments(s string) int {
	var bs, r = []byte{}, 0
	for _, c := range []byte(s) {
		if c == ' ' {
			if len(bs) > 0 {
				r += 1
				bs = []byte{}
			}
			continue
		}
		bs = append(bs, c)
	}
	if len(bs) > 0 {
		r += 1
	}
	return r
}
