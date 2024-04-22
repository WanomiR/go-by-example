package sorting

import (
	"reflect"
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
