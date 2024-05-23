package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	type test struct {
		input []string
		err   error
		want  [][]int
	}

	tests := []test{
		{
			input: []string{"*/15", "0", "1,15", "*", "1-5"},
			want: [][]int{
				{0, 15, 30, 45},
				{0},
				{1, 15},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				{1, 2, 3, 4, 5},
			},
			err: nil,
		},
		{
			input: []string{"4/15", "0,10,13", "1-8", "*", "1-5"},
			want: [][]int{
				{4, 19, 34, 49},
				{0, 10, 13},
				{1, 2, 3, 4, 5, 6, 7, 8},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				{1, 2, 3, 4, 5},
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		for i, part := range tc.input {
			actual, err := parsePart(PartType(i), part)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.want[i], actual)
		}
	}
}
