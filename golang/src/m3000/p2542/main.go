package p2542

// Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
// Output: 12
// Explanation:
// The four possible subsequence scores are:
// - We choose the indices 0, 1, and 2 with score = (1+3+3) * min(2,1,3) = 7.
// - We choose the indices 0, 1, and 3 with score = (1+3+2) * min(2,1,4) = 6.
// - We choose the indices 0, 2, and 3 with score = (1+3+2) * min(2,3,4) = 12.
// - We choose the indices 1, 2, and 3 with score = (3+3+2) * min(1,3,4) = 8.
// Therefore, we return the max score, which is 12.

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)

	var dp = make([][]int64, k)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
}
