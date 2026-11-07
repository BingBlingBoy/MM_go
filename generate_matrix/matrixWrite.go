package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
