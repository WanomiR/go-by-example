package sorting

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
