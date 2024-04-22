package sorting

import (
	"math/rand"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one value", []int{0}, []int{0}},
		{"repeating value", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negative values", []int{-4, -1, -2, -3, -5}, []int{-5, -4, -3, -2, -1}},
		{"positive values", []int{4, 1, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"values from -10 to 10", []int{10, -1, -2, 4, 3, 8, 0, -5, -6, 5}, []int{-6, -5, -2, -1, 0, 3, 4, 5, 8, 10}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.arr)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Calling BubbleSort() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	data := genData()
	var count int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
		atomic.AddInt64(&count, 1)
	}
}

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one value", []int{0}, []int{0}},
		{"repeating value", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negative values", []int{-4, -1, -2, -3, -5}, []int{-5, -4, -3, -2, -1}},
		{"positive values", []int{4, 1, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"values from -10 to 10", []int{10, -1, -2, 4, 3, 8, 0, -5, -6, 5}, []int{-6, -5, -2, -1, 0, 3, 4, 5, 8, 10}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectionSort(tt.arr)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Calling SelectionSort() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	data := genData()
	var count int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(data)
		atomic.AddInt64(&count, 1)
	}
}

func TestInsertionSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one value", []int{0}, []int{0}},
		{"repeating value", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negative values", []int{-4, -1, -2, -3, -5}, []int{-5, -4, -3, -2, -1}},
		{"positive values", []int{4, 1, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"values from -10 to 10", []int{10, -1, -2, 4, 3, 8, 0, -5, -6, 5}, []int{-6, -5, -2, -1, 0, 3, 4, 5, 8, 10}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertionSort(tt.arr)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Calling InsertionSort() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	data := genData()
	var count int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(data)
		atomic.AddInt64(&count, 1)
	}
}

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one value", []int{0}, []int{0}},
		{"repeating value", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negative values", []int{-4, -1, -2, -3, -5}, []int{-5, -4, -3, -2, -1}},
		{"positive values", []int{4, 1, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"values from -10 to 10", []int{10, -1, -2, 4, 3, 8, 0, -5, -6, 5}, []int{-6, -5, -2, -1, 0, 3, 4, 5, 8, 10}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.arr, 0, len(tt.arr)-1)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Calling QuickSort() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkQuickSortSort(b *testing.B) {
	data := genData()
	var count int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(data, 0, len(data)-1)
		atomic.AddInt64(&count, 1)
	}
}

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one value", []int{0}, []int{0}},
		{"repeating value", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negative values", []int{-4, -1, -2, -3, -5}, []int{-5, -4, -3, -2, -1}},
		{"positive values", []int{4, 1, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"values from -10 to 10", []int{10, -1, -2, 4, 3, 8, 0, -5, -6, 5}, []int{-6, -5, -2, -1, 0, 3, 4, 5, 8, 10}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.arr)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Calling MergeSort() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkMergeSortSort(b *testing.B) {
	data := genData()
	var count int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(data)
		atomic.AddInt64(&count, 1)
	}
}

func TestCountSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one value", []int{0}, []int{0}},
		{"repeating value", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negative values", []int{-4, -1, -2, -3, -5}, []int{-5, -4, -3, -2, -1}},
		{"positive values", []int{4, 1, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"values from -10 to 10", []int{10, -1, -2, 4, 3, 8, 0, -5, -6, 5}, []int{-6, -5, -2, -1, 0, 3, 4, 5, 8, 10}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := CountSort(tt.arr)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Calling CountSort() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkCountSortSort(b *testing.B) {
	data := genData()
	var count int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountSort(data)
		atomic.AddInt64(&count, 1)
	}
}

func TestHeapSort(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one value", []int{0}, []int{0}},
		{"repeating value", []int{1, 1, 1}, []int{1, 1, 1}},
		{"negative values", []int{-4, -1, -2, -3, -5}, []int{-5, -4, -3, -2, -1}},
		{"positive values", []int{4, 1, 2, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"values from -10 to 10", []int{10, -1, -2, 4, 3, 8, 0, -5, -6, 5}, []int{-6, -5, -2, -1, 0, 3, 4, 5, 8, 10}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := HeapSort(tt.arr)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Calling HeapSort() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkHeapSortSort(b *testing.B) {
	data := genData()
	var count int64

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HeapSort(data)
		atomic.AddInt64(&count, 1)
	}
}

func genData() []int {
	rand.Seed(2_147_483_647)

	n := 1_000

	data := make([]int, 0, n)
	for j := 0; j < n; j++ {
		data = append(data, rand.Intn(1_001)-500)
	}

	return data
}
