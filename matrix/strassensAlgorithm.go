package matrix

import "fmt"

type matrixView struct {
	data      []int
	rowOffset int
	colOffset int
	size      int
	stride    int
}

func (m matrixView) at(r, c int) int {
	return m.data[(m.rowOffset+r)*m.stride+(m.colOffset+c)]
}

func (m matrixView) set(r, c, val int) {
	m.data[(m.rowOffset+r)*m.stride+(m.colOffset+c)] = val
}

func (m matrixView) subView(rowDelta, colDelta, size int) matrixView {
	return matrixView{
		data:      m.data,
		rowOffset: m.rowOffset + rowDelta,
		colOffset: m.colOffset + colDelta,
		size:      size,
		stride:    m.stride,
	}
}

func (m matrixView) splitQuadrants() (tl, tr, bl, br matrixView) {
	half := m.size / 2
	return m.subView(0, 0, half),
		m.subView(0, half, half),
		m.subView(half, 0, half),
		m.subView(half, half, half)
}

func (m Matrix) asView() matrixView {
	return matrixView{
		data:      m.data,
		rowOffset: 0,
		colOffset: 0,
		size:      m.rows,
		stride:    m.cols,
	}
}

func strassenMultiply(viewA, viewB, viewC matrixView) {
	if viewA.size == 2 {
		process2x2Strassen(viewA, viewB, viewC)
		return
	}

	// TODO: Recursive case
}

func process2x2Strassen(viewA, viewB, viewC matrixView) {
	a := viewA.at(0, 0)
	b := viewA.at(0, 1)
	c := viewA.at(1, 0)
	d := viewA.at(1, 1)

	e := viewB.at(0, 0)
	f := viewB.at(0, 1)
	g := viewB.at(1, 0)
	h := viewB.at(1, 1)

	p1 := a * (f - h)
	p2 := (a + b) * h
	p3 := (c + d) * e
	p4 := d * (g - e)
	p5 := (a + d) * (e + h)
	p6 := (b - d) * (g + h)
	p7 := (a - c) * (e + f)

	viewC.set(0, 0, p5+p4-p2+p6)
	viewC.set(0, 1, p1+p2)
	viewC.set(1, 0, p3+p4)
	viewC.set(1, 1, p1+p5-p3-p7)
}

func StrassensAlgorithm(matrices []Matrix) (Matrix, error) {
	if len(matrices) != 2 {
		return Matrix{}, fmt.Errorf("expected 2 matrices, got %d", len(matrices))
	}

	a, b := matrices[0], matrices[1]
	if a.rows != a.cols || b.rows != b.cols || a.cols != b.rows {
		return Matrix{}, fmt.Errorf("matrices must be square and of equal size")
	}
	if a.rows < 2 || (a.rows&(a.rows-1)) != 0 {
		return Matrix{}, fmt.Errorf("matrix size %d must be a power of 2 >= 2", a.rows)
	}

	out := Matrix{
		rows: a.rows,
		cols: b.cols,
		data: make([]int, a.rows*b.cols),
	}

	strassenMultiply(a.asView(), b.asView(), out.asView())

	return out, nil
}
