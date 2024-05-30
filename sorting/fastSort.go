package sorting

func FastSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]

	var less, greater []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = FastSort(less)
	greater = FastSort(greater)

	return append(append(less, pivot), greater...)
}
