package search

/*Бинарный поиск. Возвращает -1 если слайс с нулевой длиной или элемента нет*/
func BinarySearch(a []int, b int) int {
	if len(a) == 0 || b > len(a)-1 || b < 0 {
		return -1
	}
	var middle int = len(a) / 2
	for {
		if a[middle] == b {
			return middle
		}
		if (middle == len(a)) || (middle == 0) {
			return -1
		}
		if a[middle] < b {
			middle = (len(a) + middle) / 2
		}
		if a[middle] > b {
			middle = middle / 2
		}
	}
}
