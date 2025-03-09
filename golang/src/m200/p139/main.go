package p139

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	} else if len(wordDict) == 0 {
		return false
	}
	n := len(s)
	var dp = make([]bool, n+1)
	dp[0] = true

	for i := range n + 1 {
		for _, word := range wordDict {
			wordLen := len(word)
			if i >= wordLen {
				if dp[i-wordLen] && s[i-wordLen:i] == word {
					dp[i] = dp[i-wordLen]
				}
			}
		}
	}
	return dp[n]
}
