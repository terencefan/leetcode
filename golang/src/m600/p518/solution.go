package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := 0; i <= amount; i++ {
			if coin <= i {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}

func main() {
	change(5, []int{1, 2, 5})
}
