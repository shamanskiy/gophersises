package testutils

import "testing"

func TestSameElements(t *testing.T) {
	sliceA := []int{1, 2, 3}
	sliceB := []int{1, 3, 2}
	sliceC := []int{1, 2}

	if !SameElements(sliceA, sliceB) {
		t.Fatalf("Error: slices have the same elements!\n")
	}

	if SameElements(sliceA, sliceC) {
		t.Fatalf("Error: slices have different elements!\n")
	}
}
