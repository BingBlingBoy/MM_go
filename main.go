package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func AddMatrix(N, M int, scanner *bufio.Scanner) [][]int {
	mat := make([][]int, N)
	for i := range mat {
		mat[i] = make([]int, M)
	}

	for i := range N {
		if !scanner.Scan() {
			log.Fatalf("expected %d rows, but file ended early at row %d", N, i)
		}

		rowLine := scanner.Text()
		rowElements := strings.Fields(rowLine)
		if len(rowElements) < M {
			log.Fatalf("row %d has fewer columns than specified by M=%d", i, M)
		}

		for j := range M {
			val, err := strconv.Atoi(rowElements[j])
			if err != nil {
				log.Fatalf("failed to convert element [%d][%d] (%q) to int: %s", i, j, rowElements[j], err)
			}
			mat[i][j] = val
		}
	}

	return mat
}

func parseMetadata(line string) (int, int, error) {
	var n, m int

	_, err := fmt.Sscanf(line, "N=%d, M=%d", &n, &m)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid matrix header format: %w", err)
	}

	return n, m, nil
}

func MultiplyMatrix(matrices [][][]int) ([][]int, error) {

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
		return nil, fmt.Errorf("cannot multiply matrice: inner dimensions mismatch (%dx%d and %dx%d)", resN, resM, resQ, resR)
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

func main() {
	filePtr := flag.String("f", "", "filename")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrices [][][]int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(line, "N") {
			n, m, err := parseMetadata(line)
			if err != nil {
				log.Fatalf("error parsing metadata: %s", err)
			}

			mat := AddMatrix(n, m, scanner)
			matrices = append(matrices, mat)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	res, err := MultiplyMatrix(matrices)
	if err != nil {
		log.Fatalf("error multiplying matrix: %s", err)
	}
	fmt.Printf("%v\n", res)
}
