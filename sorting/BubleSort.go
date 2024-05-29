package sorting

/*Сортировка пузырьком*/
func SortBuble(a []int) []int {
	n := make([]int, len(a), cap(a))
	copy(n, a)
	for i := 0; i < len(n)-1; i++ {
		for j := 0; j < len(n); j++ {
			if len(n)-1 > j {
				if n[j] > n[j+1] {
					n[j+1], n[j] = n[j], n[j+1]
				}
			}
		}
	}
	return n
}
