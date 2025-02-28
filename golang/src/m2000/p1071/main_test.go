package p1071

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
		{"", "", ""},
	}

	for _, tt := range tests {
		result := gcdOfStrings(tt.str1, tt.str2)
		if result != tt.expected {
			t.Errorf("gcdOfStrings(%v, %v) = %v; expected %v", tt.str1, tt.str2, result, tt.expected)
		}
	}
}
