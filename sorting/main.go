package main

import (
	"fmt"
	"math/rand"
	"sorting/sorting"
)

func main() {
	arr := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(21)-10)
	}
	fmt.Println(arr)

	sortedArr := sorting.MergeSort(arr)

	fmt.Println(sortedArr)
}

func selectionSort(arr []int) []int {

	n := len(arr)

	if n < 2 {
		return arr
	}

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	return arr
}
