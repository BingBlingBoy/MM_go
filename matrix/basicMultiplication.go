package matrix

import "fmt"

func BasicMultiplication(matrices []Matrix) (Matrix, error) {

	if len(matrices) < 2 {
		return Matrix{}, fmt.Errorf("matrix multiplication requires at least two matrices. Recieved %d", len(matrices))
	}

	matOne := matrices[0]
	matTwo := matrices[1]

	if len(matOne.data) == 0 || len(matTwo.data) == 0 {
		return Matrix{}, fmt.Errorf("matrices cannot contain empty dimensions")
	}

	resN, resM := matOne.rows, matOne.cols
	resQ, resR := matTwo.cols, matTwo.rows

	if resM != resR {
		return Matrix{}, fmt.Errorf("cannot multiply matrices: inner dimensions mismatch (%dx%d and %dx%d)", resN, resM, resQ, resR)
	}

	resMat := Matrix{
		rows: resN,
		cols: resQ,
		data: make([]int, resN*resQ),
	}

	for i := range resN {
		for j := range resQ {
			for k := range resM {
				val := matOne.At(i, k) * matTwo.At(k, j)
				resMat.Add(i, j, val)
			}
		}
	}

	return resMat, nil
}
