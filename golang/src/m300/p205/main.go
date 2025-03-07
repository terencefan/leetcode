package p205

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var mapping = make(map[byte]byte)
	var reverse = make(map[byte]byte)

	for i := range s {
		from, to := s[i], t[i]

		if v, ok := mapping[to]; ok {
			if v != from {
				return false
			}
		}
		mapping[to] = from

		if v, ok := reverse[from]; ok {
			if v != to {
				return false
			}
		}
		reverse[from] = to
	}
	return true
}
