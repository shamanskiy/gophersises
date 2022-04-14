package base

import "testing"

func ReportDifferentSlices[T any](got, want []T, message string, t *testing.T) {
	t.Helper()
	t.Logf("Got:\n")
	for _, elem := range got {
		t.Logf("\t%+v\n", elem)
	}
	t.Logf("Want:\n")
	for _, elem := range want {
		t.Logf("\t%+v\n", elem)
	}
	t.Errorf("%s\n", message)
}

func CheckError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Failed with error: %s\n", err)
	}
}

func SameElements[T comparable](got, want []T) bool {
	if len(got) != len(want) {
		return false
	}

	// Build hash map of occurrences in got
	diff := map[T]int{}
	for _, gotValue := range got {
		diff[gotValue]++
	}

	// Check that occurrences in want are the same
	for _, wantValue := range want {
		if _, found := diff[wantValue]; !found {
			return false
		}

		diff[wantValue]--
		if diff[wantValue] == 0 {
			delete(diff, wantValue)
		}
	}

	return true
}
