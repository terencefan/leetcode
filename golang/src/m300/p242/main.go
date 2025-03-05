package p242

func isAnagram(s string, t string) bool {
	var r = make([]int, 26)
	for _, c := range s {
		r[c-'a']++
	}
	for _, c := range t {
		r[c-'a']--
	}
	for _, v := range r {
		if v != 0 {
			return false
		}
	}
	return true
}
