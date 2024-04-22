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

	sortedArr := sorting.QuickSort(arr, 0, len(arr)-1)

	fmt.Println(sortedArr)
}
