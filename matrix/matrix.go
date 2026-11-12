package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMatrix(rows []string, n, m int) ([][]int, error) {
	if len(rows) < n {
		return nil, fmt.Errorf("expected %d rows, but only got %d", n, len(rows))
	}

	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)

		rowElements := strings.Fields(rows[i])
		if len(rowElements) < m {
			return nil, fmt.Errorf("row %d has fewer columns than specified by M=%d", i, m)
		}

		for j := range m {
			val, err := strconv.Atoi(rowElements[j])
			if err != nil {
				return nil, fmt.Errorf("failed to convert element [%d][%d] (%q) to int: %w", i, j, rowElements[j], err)
			}
			mat[i][j] = val
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
