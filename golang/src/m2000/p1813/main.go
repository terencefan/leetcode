package p1813

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence1) < len(sentence2) {
		sentence1, sentence2 = sentence2, sentence1
	}
	l1, l2 := len(sentence1), len(sentence2)

	i, j, checkpoint := 0, len(sentence2)-1, 0

	for i < len(sentence2) {
		if sentence1[i] != sentence2[i] {
			break
		}
		if sentence1[i] == ' ' {
			checkpoint = i
		}
		i++
	}

	if i == len(sentence2) {
		if l1 == l2 {
			return true
		} else if i < len(sentence1) && sentence1[i] == ' ' {
			return true
		}
	}
	i = checkpoint

	for j >= i {
		if sentence1[l1-l2+j] != sentence2[j] {
			return false
		}
		j--
	}
	return (j == -1 && sentence1[l1-l2+j+1] == ' ') || sentence2[j+1] == ' '
}
