package p71

func slash(s *[][]byte, bytes []byte) {
	if len(bytes) == 1 && bytes[0] == '.' {
		// do nothing
	} else if len(bytes) == 2 && bytes[0] == '.' && bytes[1] == '.' {
		if len(*s) > 0 {
			*s = (*s)[:len(*s)-1]
		}
	} else if len(bytes) > 0 {
		*s = append(*s, bytes)
	}
}

func simplifyPath(path string) string {
	if path[0] != '/' {
		return "/"
	}

	var stack [][]byte
	var bytes = make([]byte, 0)

	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			slash(&stack, bytes)
			bytes = make([]byte, 0)
		} else {
			bytes = append(bytes, path[i])
		}
	}

	if len(bytes) > 0 {
		slash(&stack, bytes)
	}

	if len(stack) == 0 {
		return "/"
	}

	var r = make([]byte, 0)
	for _, bytes := range stack {
		r = append(r, '/')
		r = append(r, bytes...)
	}
	return string(r)
}
