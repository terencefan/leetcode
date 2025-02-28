package p215

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{5, 2, 4, 1, 3, 6, 0}, 4, 3},
		{[]int{1}, 1, 1},
		{[]int{2, 1}, 2, 1},
	}

	for _, tt := range tests {
		result := findKthLargest(tt.nums, tt.k)
		if result != tt.expected {
			t.Errorf("findKthLargest(%v, %d) = %d; expected %d", tt.nums, tt.k, result, tt.expected)
		}
	}
}
