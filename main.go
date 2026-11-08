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

	for i := 0; i < N; i++ {
		if !scanner.Scan() {
			log.Fatalf("expected %d rows, but file ended early at row %d", N, i)
		}

		rowLine := scanner.Text()
		rowElements := strings.Fields(rowLine)
		if len(rowElements) < M {
			log.Fatalf("row %d has fewer columns than specified by M=%d", i, M)
		}

		for j := 0; j < M; j++ {
			val, err := strconv.Atoi(rowElements[j])
			if err != nil {
				log.Fatalf("failed to convert element [%d][%d] (%q) to int: %s", i, j, rowElements[j], err)
			}
			mat[i][j] = val
		}
	}

	return mat
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
	var N, M int
	var matOne [][]int
	var matTwo [][]int
	twoYet := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Contains(line, "N") {
			matInfo := strings.Split(line, ",")
			if len(matInfo) < 2 {
				log.Fatalf("invalid metadata line: %s", line)
			}

			cleanN := strings.ReplaceAll(matInfo[0], " ", "")
			cleanN = strings.ReplaceAll(cleanN, "N=", "")
			N, err = strconv.Atoi(cleanN)
			if err != nil {
				log.Fatalf("failed to parse N: %s", err)
			}

			cleanM := strings.ReplaceAll(matInfo[1], " ", "")
			cleanM = strings.ReplaceAll(cleanM, "M=", "")
			M, err = strconv.Atoi(cleanM)
			if err != nil {
				log.Fatalf("failed to parse M: %s", err)
			}

			if !twoYet {
				matOne = AddMatrix(N, M, scanner)
				twoYet = true
			} else {
				matTwo = AddMatrix(N, M, scanner)
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Println(matOne)
	fmt.Println(matTwo)
}
