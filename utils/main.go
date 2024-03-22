package main

import (
	"fmt"
	"math/rand"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func main() {
	sliceInt := generateSliceInt(42, 5)

	fmt.Println(sliceInt)
	fmt.Println(Average(sliceInt))
	fmt.Println(Sum(sliceInt))
	fmt.Println(Reverse(sliceInt))

	sliceFloat := generateSliceFloat(42, 5)
	fmt.Println(sliceFloat)
	fmt.Println(Average(sliceFloat))
	fmt.Println(Sum(sliceFloat))
	fmt.Println(Reverse(sliceFloat))

}

func Average[T Number](slice []T) (avg float64) {
	for _, v := range slice {
		avg += float64(v)
	}
	return avg / float64(len(slice))
}

func Sum[T Number](slice []T) (sum T) {
	for _, v := range slice {
		sum += v
	}
	return
}

func Reverse[T Number](slice []T) []T {
	n := len(slice)
	reversed := make([]T, 0, n)

	if n < 2 {
		return slice
	}

	for i := n - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}

func generateSliceInt(seed int64, n int) []int {
	rand.Seed(seed)
	slice := make([]int, 0, n)

	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(20))
	}
	return slice
}

func generateSliceFloat(seed int64, n int) []float64 {
	rand.Seed(seed)
	slice := make([]float64, 0, n)

	for i := 0; i < n; i++ {
		slice = append(slice, rand.NormFloat64()*1.5)
	}
	return slice
}
