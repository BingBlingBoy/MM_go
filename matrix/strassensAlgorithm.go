package matrix

// Only 2x2 currently
func StrassensAlgorithm(matrices []Matrix) (Matrix, error) {
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

	temp := make([]int, 0, 8)
	for _, mat := range matrices {
		temp = append(temp, mat.data...)
	}

	p1 := temp[a] * (temp[f] - temp[h])
	p2 := (temp[a] + temp[b]) * temp[h]
	p3 := (temp[c] + temp[d]) * temp[e]
	p4 := temp[d] * (temp[g] - temp[e])
	p5 := (temp[a] + temp[d]) * (temp[e] + temp[h])
	p6 := (temp[b] - temp[d]) * (temp[g] + temp[h])
	p7 := (temp[a] - temp[c]) * (temp[e] + temp[f])

	res := Matrix{
		rows: 2,
		cols: 2,
		data: []int{
			p5 + p4 - p2 + p6, p1 + p2,
			p3 + p4, p1 + p5 - p3 - p7,
		},
	}

	return res, nil
}
