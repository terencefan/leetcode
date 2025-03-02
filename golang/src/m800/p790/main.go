package p790

const MOD = 1e9 + 7

func numTilings(n int) int {
	if n == 0 {
		return 0
	}

	var dp = []int{1, 1, 2}
	var partial = []int{0, 0, 2}

	if n < 3 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		dp = append(dp, (dp[i-1]+dp[i-2]+2*partial[i-1])%MOD)
		partial = append(partial, (partial[i-1]+dp[i-2])*MOD)
	}
	return dp[n]
}
