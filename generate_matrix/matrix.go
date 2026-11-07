package main

import "math/rand/v2"

func generateMatrix(n, m, min, max int) [][]int {
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			matrix[i][j] = rand.IntN(max-min) + min
		}
	}

	return matrix
}
