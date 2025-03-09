package p68

import "strings"

func fill(words []string, spaces int, last bool) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaces)
	}

	var r = ""
	blanks := spaces / (len(words) - 1)
	q := spaces % (len(words) - 1)

	for i, word := range words {
		r += word

		if i < len(words)-1 {
			if last {
				r += " "
				spaces -= 1
			} else if q > 0 {
				r += strings.Repeat(" ", blanks+1)
				q--
			} else {
				r += strings.Repeat(" ", blanks)
			}
		}
	}
	if last {
		r += strings.Repeat(" ", spaces)
	}
	return r
}

func fullJustify(words []string, maxWidth int) []string {

	var r = make([]string, 0)
	var temp = make([]string, 0)
	var length = maxWidth

	for _, word := range words {
		if len(word) > length-len(temp) {
			r = append(r, fill(temp, length, false))
			temp = temp[:1]
			temp[0] = word
			length = maxWidth - len(word)
		} else {
			temp = append(temp, word)
			length -= len(word)
		}
	}
	if len(temp) > 0 {
		r = append(r, fill(temp, length, true))
	}
	return r
}
