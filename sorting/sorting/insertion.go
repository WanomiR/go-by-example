package sorting

import "fmt"

func InsertionSort(arr []int) []int {
	n := len(arr)

	if n < 2 {
		return arr
	}

	for i := 1; i < n; i++ {
		j := i - 1
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		fmt.Println(j)
		arr[j+1] = key
	}

	return arr
}
