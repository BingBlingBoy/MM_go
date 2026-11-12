package matrix

// Only 2x2 currently
func StrassensAlgorithm(matrices [][][]int) ([][]int, error) {

	const (
		a int = iota
		b
		c
		d
		e
		f
		g
		h
	)

	tempMat := make([]int, 0, 8)

	for _, mat := range matrices {
		for i := range mat {
			for j := range mat[i] {
				tempMat = append(tempMat, mat[i][j])
			}
		}
	}

	p1 := tempMat[a] * (tempMat[f] - tempMat[h])
	p2 := (tempMat[a] + tempMat[b]) * tempMat[h]
	p3 := (tempMat[c] + tempMat[d]) * tempMat[e]
	p4 := tempMat[d] * (tempMat[g] - tempMat[e])
	p5 := (tempMat[a] + tempMat[d]) * (tempMat[e] + tempMat[h])
	p6 := (tempMat[b] - tempMat[d]) * (tempMat[g] + tempMat[h])
	p7 := (tempMat[a] - tempMat[c]) * (tempMat[e] + tempMat[f])

	res := [][]int{
		{p5 + p4 - p2 + p6, p1 + p2},
		{p3 + p4, p1 + p5 - p3 - p7},
	}

	return res, nil
}
