package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func writeMatrixToTxtFile(filename string, matrix [][]int, delimeter string) error {
	_, statErr := os.Stat(filename)
	fileExists := statErr == nil

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to create/open file: %s", err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	if fileExists {
		_, err := w.WriteString("\n")
		if err != nil {
			return fmt.Errorf("failed to write separator newline: %w", err)
		}
	}

	for _, row := range matrix {
		for j, val := range row {
			_, err := w.WriteString(strconv.Itoa(val))
			if err != nil {
				return fmt.Errorf("failed to write value: %s", err)
			}

			if j < len(row)-1 {
				_, err := w.WriteString(delimeter)
				if err != nil {
					return fmt.Errorf("failed to write delimeter: %s", err)
				}
			}
		}

		_, err = w.WriteString("\n")
		if err != nil {
			return fmt.Errorf("failed to write newline: %s", err)
		}
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("failed to flush buffer: %s", err)
	}
	return nil
}

func main() {
	nPtr := flag.Int("n", 0, "n dimension")
	mPtr := flag.Int("m", 0, "m dimension")
	rangePtr := flag.String("range", "", "range of values e.g. a,b")
	filePtr := flag.String("f", "", "text file")

	flag.Parse()

	n := *nPtr
	m := *mPtr

	splitRange := strings.Split(*rangePtr, ",")

	lowerBound, err := strconv.Atoi(splitRange[0])
	if err != nil {
		log.Fatalf("failed to convert to int: %s", err)
	}
	upperBound, err := strconv.Atoi(splitRange[1])
	if err != nil {
		log.Fatalf("failed to convert to int: %s", err)
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			matrix[i][j] = randRange(lowerBound, upperBound)
		}
	}

	err = writeMatrixToTxtFile(*filePtr, matrix, " ")
	if err != nil {
		log.Fatalf("failed to write to matrix: %s", err)
	}

	fmt.Println(matrix)
}
