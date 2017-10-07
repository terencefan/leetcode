package p316

import "ds"

func removeDuplicateLetters(s string) string {
	var a = [256]int{}
	var b = [256]bool{}

	for _, x := range s {
		a[x]++
	}

	var st = ds.NewStack(byte('_'))

	for _, x := range s {
		a[x]--
		if b[x] {
			continue
		}
		var back = st.Back().(byte)
		for byte(x) < back && a[back] > 0 {
			b[back] = false
			st.Pop()
			back = st.Back().(byte)
		}
		st.Push(byte(x))
		b[x] = true
	}

	var res = make([]byte, 32)
	for _, c := range st.All()[1:] {
		res = append(res, c.(byte))
	}
	return string(res)
}

func RemoveDuplicateLetters(s string) string {
	return removeDuplicateLetters(s)
}
