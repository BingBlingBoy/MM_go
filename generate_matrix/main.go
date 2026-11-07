package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Config struct {
	n          int
	m          int
	lowerBound int
	upperBound int
	filename   string
}

func main() {
	cfg, err := ParseAndValidateFlags()
	if err != nil {
		log.Fatalf("failed to parse and validate flags: %s", err)
	}

	matrix := generateMatrix(cfg.n, cfg.m, cfg.lowerBound, cfg.upperBound)

	err = writeMatrixToTxtFile(cfg.filename, matrix, " ")
	if err != nil {
		log.Fatalf("failed to write to matrix: %s", err)
	}

	fmt.Println(matrix)
}

func ParseAndValidateFlags() (Config, error) {
	nPtr := flag.Int("n", 0, "n dimension")
	mPtr := flag.Int("m", 0, "m dimension")
	rangePtr := flag.String("range", "", "range of values e.g. a,b")
	filePtr := flag.String("f", "", "text file")

	flag.Parse()

	if *nPtr <= 0 || *mPtr <= 0 {
		return Config{}, fmt.Errorf("-n and -m dimensions must be integers greater than 0")
	}
	if *filePtr == "" {
		return Config{}, fmt.Errorf("target file path (-f) is required")
	}
	if *rangePtr == "" {
		return Config{}, fmt.Errorf("value range (-range) is required")
	}

	splitRange := strings.Split(*rangePtr, ",")
	if len(splitRange) != 2 {
		return Config{}, fmt.Errorf("invalid range format; expected 'min,max' (e.g. 5,20)")
	}

	lowerBound, err := strconv.Atoi(splitRange[0])
	if err != nil {
		return Config{}, fmt.Errorf("invalid minimum range value %q: %w", splitRange[0], err)
	}

	upperBound, err := strconv.Atoi(splitRange[1])
	if err != nil {
		return Config{}, fmt.Errorf("invalid maximum range value %q: %w", splitRange[1], err)
	}

	if lowerBound >= upperBound {
		return Config{}, fmt.Errorf("minimum range bound (%d) must be strictly less than maximum bound (%d)", lowerBound, upperBound)
	}

	return Config{
		n:          *nPtr,
		m:          *mPtr,
		lowerBound: lowerBound,
		upperBound: upperBound,
		filename:   *filePtr,
	}, err
}
