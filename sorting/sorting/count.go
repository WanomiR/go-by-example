package sorting

func CountSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	min, delta := findMin(arr), 0
	if min < 0 {
		delta = 0 - min
		for i := 0; i < len(arr); i++ {
			arr[i] += delta
		}
	}

	max := findMax(arr)
	count := make([]int, max+1)
	sortedArr := make([]int, len(arr))

	for _, num := range arr {
		count[num]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		sortedArr[count[num]-1] = num
		count[num]--
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = sortedArr[i] - delta
	}

	return arr
}

func findMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func findMin(arr []int) int {
	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min
}
