package main

import "fmt"

func main() {

	var n int

	scanned, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	} else if scanned != 1 || n < 0 || n > 30 {
		err = fmt.Errorf("pass in a positive integer not greater than 30")
		panic(err)
	}

	fmt.Printf("Factorial    %d is: %d\n", n, factorial(n))
	fmt.Printf("Subfactorial %d is: %d", n, subFactorial(n))
}

func factorial(n int) uint64 {
	if n < 2 {
		return 1
	}
	return uint64(n) * factorial(n-1)
}

func subFactorial(n int) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}

	var a int64 = -1
	if n%2 == 0 {
		a = 1
	}
	return int64(n)*subFactorial(n-1) + a
}
