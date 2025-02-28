package p841

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		rooms    [][]int
		expected bool
	}{
		{
			rooms:    [][]int{{1}, {2}, {3}, {}},
			expected: true,
		},
		{
			rooms:    [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			expected: false,
		},
		{
			rooms:    [][]int{{1, 2, 3}, {0}, {0}, {0}},
			expected: true,
		},
		{
			rooms:    [][]int{{1}, {2}, {0}, {}},
			expected: true,
		},
		{
			rooms:    [][]int{{1}, {}, {2}, {3}},
			expected: false,
		},
	}

	for _, tt := range tests {
		result := canVisitAllRooms(tt.rooms)
		if result != tt.expected {
			t.Errorf("canVisitAllRooms(%v) = %v; expected %v", tt.rooms, result, tt.expected)
		}
	}
}
