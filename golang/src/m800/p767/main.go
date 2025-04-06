package p767

func reorganizeString(s string) string {
	var arr [26]int

	for _, c := range s {
		arr[c-'a']++
	}

	b, maxb := 0, 0
	for i := range 26 {
		if arr[i] > maxb {
			maxb = arr[i]
			b = i
		}
	}

	if maxb > (len(s)+1)/2 {
		return ""
	}

	var r = make([]byte, len(s))
	i, k := 0, 0
	for ; i < len(s); i += 2 {
		if arr[b] > 0 {
			arr[b]--
			r[i] = byte(b + 'a')
			continue
		}
		for k < 26 && arr[k] == 0 {
			k++
		}
		arr[k]--
		r[i] = byte(k + 'a')
	}

	for i = 1; i < len(s); i += 2 {
		for k < 26 && arr[k] == 0 {
			k++
		}
		arr[k]--
		r[i] = byte(k + 'a')
	}
	return string(r)
}
