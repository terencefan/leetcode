package p17

var dmap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func dfs(digits string, index int, path []rune, res *[]string) {
	if index == len(digits) {
		*res = append(*res, string(path))
		return
	}

	if runes, ok := dmap[digits[index]]; ok {
		for _, r := range runes {
			path = append(path, r)
			dfs(digits, index+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var res = make([]string, 0)
	var path = make([]rune, 0)
	dfs(digits, 0, path, &res)
	return res
}
