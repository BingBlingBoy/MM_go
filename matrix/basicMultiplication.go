package matrix

import "fmt"

func BasicMultiplication(matrices [][][]int) ([][]int, error) {

	if len(matrices) < 2 {
		return nil, fmt.Errorf("matrix multiplication requires at least two matrices. Recieved %d", len(matrices))
	}

	matOne := matrices[0]
	matTwo := matrices[1]

	if len(matOne) == 0 || len(matOne[0]) == 0 || len(matTwo) == 0 || len(matTwo[0]) == 0 {
		return nil, fmt.Errorf("matrices cannot contain empty dimensions")
	}

	resN, resM := len(matOne), len(matOne[0])
	resQ, resR := len(matTwo[0]), len(matTwo)

	if resM != resR {
		return nil, fmt.Errorf("cannot multiply matrices: inner dimensions mismatch (%dx%d and %dx%d)", resN, resM, resQ, resR)
	}

	resMat := make([][]int, resN)
	for i := range resMat {
		resMat[i] = make([]int, resQ)
	}

	for i := range resN {
		for j := range resQ {
			for k := range resM {
				resMat[i][j] += matOne[i][k] * matTwo[k][j]
			}
		}
	}

	return resMat, nil
}
