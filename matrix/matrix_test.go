package matrix

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"strings"
	"testing"
)

var benchResult Matrix

func RandRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func NewRandomMatrix(size int) []int {
	mat := make([]int, size*size)

	for i := range size {
		for j := range size {
			mat[i*size+j] = RandRange(-10000, 10000)
		}
	}

	return mat
}

func BenchmarkBasicMatMult(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			matrices := []Matrix{
				{
					rows: size,
					cols: size,
					data: NewRandomMatrix(size),
				},
				{
					rows: size,
					cols: size,
					data: NewRandomMatrix(size),
				},
			}

			var err error
			b.ReportAllocs()
			b.ResetTimer()

			for b.Loop() {
				benchResult, err = BasicMultiplication(matrices)
				if err != nil {
					b.Fatal(err)
				}
			}

			msPerOp := float64(b.Elapsed().Nanoseconds()) / 1e6 / float64(b.N)
			b.ReportMetric(msPerOp, "ms/op")
		})
	}
}

func TestBasicMatMult(t *testing.T) {
	tests := []struct {
		name        string
		matrices    []Matrix
		want        Matrix
		wantErr     bool
		errContains string
	}{
		{
			name: "Successful 2x2 multiplication",
			matrices: []Matrix{
				{
					rows: 2,
					cols: 2,
					data: []int{1, 2, 3, 4},
				},
				{
					rows: 2,
					cols: 2,
					data: []int{5, 6, 7, 8},
				},
			},
			want: Matrix{
				rows: 2,
				cols: 2,
				data: []int{19, 22, 43, 50},
			},
			wantErr: false,
		},
		{
			name: "Error: less than 2 matrices passed",
			matrices: []Matrix{
				{
					rows: 2,
					cols: 2,
					data: []int{19, 22, 43, 50},
				},
			},
			want:    Matrix{},
			wantErr: true,
		},
		{
			name: "Error: matrices cannot contain empty dimensions",
			matrices: []Matrix{
				{
					rows: 0,
					cols: 0,
					data: []int{},
				},
				{
					rows: 0,
					cols: 0,
					data: []int{},
				},
			},
			want:    Matrix{},
			wantErr: true,
		},
		{
			name: "Error: inner dimension mismatch",
			matrices: []Matrix{
				{
					rows: 3,
					cols: 2,
					data: []int{2, 4, 1, 4, 3, 4},
				},
				{
					rows: 1,
					cols: 2,
					data: []int{2, 1},
				},
			},
			want:    Matrix{},
			wantErr: true,
		},
		{
			name: "Multiplication with negative values, zeroes, and larger values",
			matrices: []Matrix{
				{
					rows: 2,
					cols: 2,
					data: []int{0, -2, 10, 5},
				},
				{
					rows: 2,
					cols: 2,
					data: []int{3, 4, -1, 0},
				},
			},
			want: Matrix{
				rows: 2,
				cols: 2,
				data: []int{2, 0, 25, 40},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BasicMultiplication(tt.matrices)

			if (err != nil) != tt.wantErr {
				t.Fatalf("BasicMultiplication() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
				t.Errorf("BasicMultiplication() error = %q, want it to contain %q", err.Error(), tt.errContains)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicMultiplication() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestParseMatrix(t *testing.T) {
	tests := []struct {
		name    string
		row     []string
		n       int
		m       int
		want    Matrix
		wantErr bool
	}{
		{
			name: "Successful parse",
			row:  []string{"1 2", "3 4"},
			n:    2,
			m:    2,
			want: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			wantErr: false,
		},
		{
			name:    "Error: fewer columns specified by M",
			row:     []string{"1 2", "3 4"},
			n:       2,
			m:       3,
			want:    Matrix{},
			wantErr: true,
		},
		{
			name:    "Error: fewer rows provided than N asks for",
			row:     []string{"1 2"},
			n:       2,
			m:       2,
			want:    Matrix{},
			wantErr: true,
		},
		{
			name:    "Error: bad non-integer data string",
			row:     []string{"1 broken", "3 4"},
			n:       2,
			m:       2,
			want:    Matrix{},
			wantErr: true,
		},
		{
			name: "Successful parse with column truncation",
			row:  []string{"1 2 999", "3 4 888"},
			n:    2,
			m:    2,
			want: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMatrix(tt.row, tt.n, tt.m)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMatrix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMetadata(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		wantN   int
		wantM   int
		wantErr bool
	}{
		{
			name:    "Successful Parsing of Metadata",
			line:    "N=3, M=2",
			wantN:   3,
			wantM:   2,
			wantErr: false,
		},
		{
			name:    "Error: Bad formatting keywords",
			line:    "X=3, Y=2",
			wantN:   0,
			wantM:   0,
			wantErr: true,
		},
		{
			name:    "Error: Missing value elements",
			line:    "N=3, M=",
			wantN:   0,
			wantM:   0,
			wantErr: true,
		},
		{
			name:    "Validation Warning: Negative dimensions parsed",
			line:    "N=-3, M=-2",
			wantN:   -3,
			wantM:   -2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotM, err := ParseMetadata(tt.line)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}

			if gotN != tt.wantN || gotM != tt.wantM {
				t.Errorf("ParseMetadata() gotN = %d, gotM = %d, wantN %d, wantM %d", gotN, gotM, tt.wantN, tt.wantM)
			}
		})
	}
}
