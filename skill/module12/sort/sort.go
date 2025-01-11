package sort

func BubbleSort(ar []int) {

	length := len(ar)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}

	}

}
func SelectionSort(ar []int) {
	n := len(ar)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
		}
	}
}

func InsertionSort(ar []int) {
	n := len(ar)
	for i := 1; i < n; i++ {
		key := ar[i]
		j := i - 1
		for j >= 0 && ar[j] > key {
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1] = key
	}

}
func MergeSort(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}
	mid := len(ar) / 2
	left := MergeSort(ar[:mid])
	right := MergeSort(ar[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}
func partition(ar []int, low, high int) int {
	pivot := ar[high]
	i := low - 1
	for j := low; j < high; j++ {
		if ar[j] <= pivot {
			i++
			ar[i], ar[j] = ar[j], ar[i]
		}
	}
	ar[i+1], ar[high] = ar[high], ar[i+1]
	return i + 1
}

func QuickSort(ar []int, low, high int) {
	if low < high {
		pi := partition(ar, low, high)
		QuickSort(ar, low, pi-1)
		QuickSort(ar, pi+1, high)
	}
}
