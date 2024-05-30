package sorting

import "testing"

func TestFastSort(t *testing.T) {
	options := [][]int{
		{111, 3, 8, 5, 1, 0}, {0, 1, 3, 5, 8, 111},
		{1, 3}, {1, 3},
		{3, 3}, {3, 3},
		{3, 1}, {1, 3},
		{100000, 10000, 1000}, {1000, 10000, 100000},
	}
	for i := 0; i < len(options); i += 2 {
		var sum int
		res := FastSort(options[i])
		for k, val := range options[i+1] {
			if val == res[k] {
				sum++
			}
		}
		if sum != len(options[i+1]) {
			t.Errorf("must: %v, have %v", options[i+1], res)
		}
	}
}
