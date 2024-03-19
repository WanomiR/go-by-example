package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	// read arguments from the command line
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Pass in N and M dimensions as arguments!")
		return
	}

	row, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Need an integer, got: ", arguments[1])
	}

	col, err := strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("Need an integer, got: ", arguments[2])
	}
	fmt.Printf("Using %dx%d arrays\n", row, col)

	// initialize two matrices with random numbers
	m1 := createMatrix(row, col)
	m2 := createMatrix(row, col)

	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)

	r1 := addMatrices(m1, m2)
	fmt.Println("addition:")
	fmt.Println(r1)

	r2 := addMatrices(m1, negativeMatrix(m2))
	fmt.Println("subtraction:")
	fmt.Println(r2)

	r3, err := multiplyMatrices(m1, m2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("multiplication")
	fmt.Println(r3)

}

func random(min, max int) float64 {
	return float64(rand.Intn(max-min) + min) // not sure why it's done here this way
}

func createMatrix(row, col int) (result [][]float64) {
	result = make([][]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[i] = append(result[i], random(-5, i*j+rand.Intn(10)))
		}
	}
	return
}

func negativeMatrix(s [][]float64) [][]float64 {
	for i, x := range s {
		for j, _ := range x {
			s[i][j] *= -1
		}
	}
	return s
}

func addMatrices(m1 [][]float64, m2 [][]float64) [][]float64 {
	result := make([][]float64, len(m1))
	for i, x := range m1 {
		for j, _ := range x {
			result[i] = append(result[i], m1[i][j]+m2[i][j])
		}
	}
	return result
}

func multiplyMatrices(m1 [][]float64, m2 [][]float64) ([][]float64, error) {
	if len(m1[0]) != len(m2) {
		return nil, errors.New("cannot multiply matrices: dimensions mismatch")
	}

	result := make([][]float64, len(m1))
	for i, _ := range m1 {
		result[i] = make([]float64, len(m2[0]))

		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return result, nil
}

func getCofactor(A [][]float64, temp [][]float64, p int, q int, n int) {
	i, j := 0, 0

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row != p && col != 1 {
				temp[i][j] = A[row][col]
				j++
				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}

// TODO continue with matrix division
//func determinant(A [][]float64, n int) float64 {
//	D := float64(0)
//	if n == 1 {
//		return A[0][0]
//	}
//
//	temp := createMatrix(n, n)
//}
