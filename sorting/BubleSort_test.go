package sorting

import "testing"

func TestSortBuble(t *testing.T) {
	sl := []int{1, 5, 2, 6, 3, 4, 9, 8, 7}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := SortBuble(sl)
	var sum int
	for i, val := range res {
		if val == want[i] {
			sum++
		}
	}
	if sum != len(want) {
		t.Errorf("Ошибка при сортировке. %v", res)
	}
}
