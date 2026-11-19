package main

import (
	"bufio"
	mat "ckdb/matrix"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Refactor all matrices to 1D matrices

type multiplyFunc func([][][]int) ([][]int, error)

func main() {
	filePtr := flag.String("f", "", "filename")

	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrices []mat.Matrix
	var n, m int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.Contains(line, "N") {
			n, m, err = mat.ParseMetadata(line)
			if err != nil {
				log.Fatalf("error parsing metadata: %s", err)
			}

			var rawRows []string
			for range n {
				if !scanner.Scan() {
					log.Fatalf("expected %d rows under header '%s', but file ended early", n, line)
				}
				rawRows = append(rawRows, scanner.Text())
			}

			parsed, err := mat.ParseMatrix(rawRows, n, m)
			if err != nil {
				log.Fatalf("failed to build matrix: %s", err)
			}
			matrices = append(matrices, parsed)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	res, err := mat.StrassensAlgorithm(matrices)
	if err != nil {
		log.Fatalf("error multiplying matrix: %s", err)
	}
	fmt.Printf("%v\n", res)
}
