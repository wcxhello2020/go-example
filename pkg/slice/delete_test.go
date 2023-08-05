package slice

import (
	"testing"
)

func TestDelete(t *testing.T) {
	// Positive test case
	src := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 4, 5}
	result, err := Delete(src, 2)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if !equalSlices(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	// Negative test case
	src = []int{1, 2, 3, 4, 5}
	_, err = Delete(src, 5)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
