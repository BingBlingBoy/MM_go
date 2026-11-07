package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
	"strings"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func main() {
	nPtr := flag.Int("n", 0, "n dimension")
	mPtr := flag.Int("m", 0, "m dimension")
	rangePtr := flag.String("range", "", "range of values e.g. a,b")
	filePtr := flag.String("f", "", "text file")

	flag.Parse()

	fmt.Println("nptr:", *nPtr)
	fmt.Println("mptr:", *mPtr)
	fmt.Println("filePtr:", *filePtr)
	fmt.Println("rangePtr:", *rangePtr)

	// file, err := os.Create(*filePtr)
	// if err != nil {
	// 	log.Fatalf("failed to create file: %s", err)
	// }
	// defer file.Close()

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

	fmt.Println(matrix)
}
