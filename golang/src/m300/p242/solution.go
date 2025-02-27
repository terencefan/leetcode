package main

func isAnagram(s string, t string) bool {
	var arr = make([]int, 26)
	for _, c := range s {
		arr[c-'a']++
	}
	for _, c := range t {
		arr[c-'a']--
	}
	for _, x := range arr {
		if x != 0 {
			return false
		}
	}
	return true
}
