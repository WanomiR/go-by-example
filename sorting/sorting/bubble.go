package sorting

func BubbleSort(arr []int) []int {
	n := len(arr)

	if n < 2 {
		return arr
	}

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
