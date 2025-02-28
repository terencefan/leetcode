package p450

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		key      int
		expected *TreeNode
	}{
		{
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
			key:      3,
			expected: &TreeNode{Val: 5, Left: nil, Right: &TreeNode{Val: 6}},
		},
		{
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}},
			key:      3,
			expected: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 6}},
		},
		{
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
			key:      5,
			expected: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: nil},
		},
		{
			root:     &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
			key:      7,
			expected: &TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
		},
	}

	for _, tt := range tests {
		result := deleteNode(tt.root, tt.key)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("deleteNode(%v, %d) = %v; expected %v", tt.root, tt.key, result, tt.expected)
		}
	}
}
