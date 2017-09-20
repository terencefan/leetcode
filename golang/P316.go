package main

func removeDuplicateLetters(s string) string {
	var a = [256]int{}
	var b = [256]bool{}

	for _, x := range s {
		a[x]++
	}

	var st = stack(byte('_'))

	for _, x := range s {
		a[x]--
		if b[x] {
			continue
		}
		var back = st.back().(byte)
		for byte(x) < back && a[back] > 0 {
			b[back] = false
			st.pop()
			back = st.back().(byte)
		}
		st.push(byte(x))
		b[x] = true
	}

	var res = make([]byte, 32)
	for _, c := range st.all()[1:] {
		res = append(res, c.(byte))
	}
	return string(res)
}
