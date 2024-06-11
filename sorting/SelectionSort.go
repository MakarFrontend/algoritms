package sorting

/*Сортивка выбором*/
func SelectionSort(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	if len(a) == 1 {
		return a
	}
	for i := 0; i < len(a); i++ {
		minIndex := i
		for k := i + 1; k < len(a); k++ {
			if a[k] < a[minIndex] {
				minIndex = k
			}
		}
		a[minIndex], a[i] = a[i], a[minIndex]
	}
	return a
}
