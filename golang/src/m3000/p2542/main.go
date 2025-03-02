package p2542

import "sort"

// Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
// Output: 12
// Explanation:
// The four possible subsequence scores are:
// - We choose the indices 0, 1, and 2 with score = (1+3+3) * min(2,1,3) = 7.
// - We choose the indices 0, 1, and 3 with score = (1+3+2) * min(2,1,4) = 6.
// - We choose the indices 0, 2, and 3 with score = (1+3+2) * min(2,3,4) = 12.
// - We choose the indices 1, 2, and 3 with score = (3+3+2) * min(1,3,4) = 8.
// Therefore, we return the max score, which is 12.

type Pair struct {
	num1 int
	num2 int
}

type MinHeap struct {
	nums []Pair
	cap  int
}

func children(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func less(a, b Pair) bool {
	return a.num1 < b.num1
}

func (h *MinHeap) swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *MinHeap) rise(i int) {
	for i > 0 {
		p := parent(i)
		if less(h.nums[p], h.nums[i]) {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *MinHeap) sink(i int) {
	for i < len(h.nums) {
		l, r := children(i)

		var k = i
		if l < len(h.nums) && less(h.nums[l], h.nums[k]) {
			k = l
		}
		if r < len(h.nums) && less(h.nums[r], h.nums[k]) {
			k = r
		}
		if i == k {
			break
		}
		h.swap(i, k)
		i = k
	}
}

func (h *MinHeap) insert(pair Pair) {
	h.nums = append(h.nums, pair)
	h.rise(len(h.nums) - 1)
}

func (h *MinHeap) replace(pair Pair) {
	if h.cap == 0 {
		return
	}
	h.nums[0] = pair
	h.sink(0)
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)

	var pairs = make([]Pair, n)
	for i := range n {
		pairs[i] = Pair{nums1[i], nums2[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].num2 > pairs[j].num2
	})

	var r, prod int64 = 0, 0
	var heap = MinHeap{make([]Pair, 0), k}
	var sum = 0
	for i := range k {
		heap.insert(pairs[i])
		sum += pairs[i].num1
	}
	r = int64(pairs[k-1].num2) * int64(sum)

	for i := k; i < n; i++ {
		min := pairs[i].num2
		sum -= heap.nums[0].num1
		sum += pairs[i].num1

		heap.replace(pairs[i])
		prod = int64(min) * int64(sum)
		if prod > r {
			r = prod
		}
	}
	return r
}
