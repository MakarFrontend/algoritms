package search

import "testing"

func TestBinarySearch(t *testing.T) {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	wantIndex := 6
	index := BinarySearch(sl, 6)
	if index != wantIndex {
		t.Errorf("Ошибка: %v вместо %v", index, wantIndex)
	}
}
