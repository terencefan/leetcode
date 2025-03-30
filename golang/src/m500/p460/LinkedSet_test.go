package p460

import (
	"testing"
)

func TestLinkedSet_Add(t *testing.T) {
	set := NewLinkedList[int]()
	set.Add(1)
	set.Add(2)
	set.Add(1) // Duplicate, should not be added again

	if len(set.m) != 2 {
		t.Errorf("Expected set size to be 2, got %d", len(set.m))
	}

	if set.head.value != 1 || set.tail.value != 2 {
		t.Errorf("Head or tail values are incorrect")
	}
}

func TestLinkedSet_Remove(t *testing.T) {
	set := NewLinkedList[int]()
	set.Add(1)
	set.Add(2)
	set.Remove(1)

	if len(set.m) != 1 {
		t.Errorf("Expected set size to be 1, got %d", len(set.m))
	}

	if set.head.value != 2 || set.tail.value != 2 {
		t.Errorf("Head or tail values are incorrect after removal")
	}

	set.Remove(2)
	if len(set.m) != 0 || set.head != nil || set.tail != nil {
		t.Errorf("Expected set to be empty after removing all elements")
	}
}

func TestLinkedSet_AddAndRemove(t *testing.T) {
	set := NewLinkedList[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Remove(2)

	if len(set.m) != 2 {
		t.Errorf("Expected set size to be 2, got %d", len(set.m))
	}

	if set.head.value != 1 || set.tail.value != 3 {
		t.Errorf("Head or tail values are incorrect after removal")
	}

	if set.head.next.value != 3 || set.tail.prev.value != 1 {
		t.Errorf("Links between nodes are incorrect after removal")
	}
}
