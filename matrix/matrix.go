package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	rows int
	cols int
	data []int
}

func (m Matrix) At(r, c int) int {
	return m.data[r*m.cols+c]
}

func (m Matrix) Set(r, c, val int) {
	m.data[r*m.cols+c] = val
}

func (m Matrix) Rows() int { return m.rows }
func (m Matrix) Cols() int { return m.cols }

func ParseMatrix(rows []string, n, m int) (Matrix, error) {
	if n <= 0 || m <= 0 {
		return Matrix{}, fmt.Errorf("invalid dimensions: %dx%d", n, m)
	}

	if len(rows) < n {
		return Matrix{}, fmt.Errorf("expected %d rows, but only got %d", n, len(rows))
	}

	mat := Matrix{rows: n, cols: m, data: make([]int, n*m)}
	for i := range n {

		rowElements := strings.Fields(rows[i])
		if len(rowElements) < m {
			return Matrix{}, fmt.Errorf("row %d has fewer columns than specified by M=%d", i, m)
		}

		for j := range m {
			val, err := strconv.Atoi(rowElements[j])
			if err != nil {
				return Matrix{}, fmt.Errorf("failed to convert element [%d][%d] (%q) to int: %w", i, j, rowElements[j], err)
			}
			mat.data[i*m+j] = val
		}
	}

	return mat, nil
}

func ParseMetadata(line string) (int, int, error) {
	var n, m int

	_, err := fmt.Sscanf(line, "N=%d, M=%d", &n, &m)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid matrix header format: %w", err)
	}

	return n, m, nil
}
