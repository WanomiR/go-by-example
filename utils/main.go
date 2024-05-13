package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Map map[string]int

func main() {
	words := []string{"hello", "world", "nice", "world", "day", "hello", "fuck off"}
	fmt.Println(countUnique(words))
}

func countUnique[T comparable](arr []T) int {
	if len(arr) < 2 {
		return len(arr)
	}

	unique := 1
	for i := 1; i < len(arr); i++ {
		j := 0
		for j < i {
			if arr[i] == arr[j] {
				break
			}
			j++
		}
		if i == j {
			unique++
		}
	}
	return unique
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

func slicesAreEqual[T Number](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}

func mapsAreEqual[T Map](map1, map2 T) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if v != map2[k] {
			return false
		}
	}
	return true
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

func heapSort[T Number](slice []T) []T {
	var heapify func(slice []T, n, i int)
	n := len(slice)

	heapify = func(slice []T, n, i int) {
		var largest, left, right int

		largest = i
		left = 2*i + 1
		right = 2*i + 2

		if (left < n) && (slice[left] > slice[largest]) {
			largest = left
		}

		if (right < n) && (slice[right] > slice[largest]) {
			largest = right
		}

		if largest != i {
			slice[i], slice[largest] = slice[largest], slice[i]

			heapify(slice, n, largest)
		}
	}

	// build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(slice, n, i)
	}

	// heap sort
	for i := n - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, i, 0)
	}

	return slice
}

func retrieveDirPath(filePath string) string {
	subdirs := strings.Split(filePath, "/")

	if len(subdirs) < 2 {
		return "./"
	}

	return strings.Join(subdirs[:len(subdirs)-1], "/")
}
