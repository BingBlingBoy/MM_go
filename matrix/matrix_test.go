package matrix

import (
	"reflect"
	"testing"
)

func TestBasicMatMult(t *testing.T) {
	tests := []struct {
		name        string
		matrices    [][][]int
		want        [][]int
		wantErr     bool
		errContains string
	}{
		{
			name: "Successful 2x2 multiplication",
			matrices: [][][]int{
				{
					{1, 2},
					{3, 4},
				},
				{
					{5, 6},
					{7, 8},
				},
			},
			want: [][]int{
				{19, 22},
				{43, 50},
			},
			wantErr: false,
		},
		{
			name: "Error: less than 2 matrices passed",
			matrices: [][][]int{
				{
					{1, 2},
					{3, 4},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error: matrices cannot contain empty dimensions",
			matrices: [][][]int{
				{
					{},
				},
				{
					{},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error: inner dimension mismatch",
			matrices: [][][]int{
				{
					{2, 4},
					{1, 4},
					{3, 4},
				},
				{
					{2, 1},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MultiplyMatrix(tt.matrices)

			if (err != nil) != tt.wantErr {
				t.Fatalf("MultiplyMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyMatrix() got = %v, want %v", got, tt.want)
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
		want    [][]int
		wantErr bool
	}{
		{
			name: "Successful parse",
			row:  []string{"1 2", "3 4"},
			n:    2,
			m:    2,
			want: [][]int{
				{1, 2},
				{3, 4},
			},
			wantErr: false,
		},
		{
			name:    "Error: fewer columns specified by M",
			row:     []string{"1 2", "3 4"},
			n:       2,
			m:       3,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Error: fewer rows provided than N asks for",
			row:     []string{"1 2"},
			n:       2,
			m:       2,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Error: bad non-integer data string",
			row:     []string{"1 broken", "3 4"},
			n:       2,
			m:       2,
			want:    nil,
			wantErr: true,
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
