package main

import (
	"fmt"
	"math/rand"
	"playground/sorting"
)

func main() {
	arr := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(21)-10)
	}
	fmt.Println(arr)

	sortedArr := sorting.SelectionSort(arr)

	fmt.Println(sortedArr)
}
