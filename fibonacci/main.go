package main

import "fmt"

func main() {
	// determine the order of a given number
	// in the fibonacci siquence

	var n, f int
	_, _ = fmt.Scan(&n)

	// initialize fibonacci function
	fib := fibonacci()

	// iterate until meet fibonacci or reach the limit
	cnt := 0
	for f < n {
		f = fib()
		if f < n {
			cnt++
		}
	}

	// print the order number if given number
	// is a fibonacci number
	if f == n {
		fmt.Print(cnt)
	} else {
		fmt.Print(-1)
	}

}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a, b = b, c+b
		return c
	}
}
