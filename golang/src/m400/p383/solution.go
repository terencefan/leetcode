package main

func canConstruct(ransomNote string, magazine string) bool {
	var m = make([]int, 256)

	for i := range magazine {
		m[magazine[i]]++
	}

	for i := range ransomNote {
		m[ransomNote[i]]--
		if m[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}
