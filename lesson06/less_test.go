package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Concat(slices [][]int) []int {
	s := []int{}

	for _, val := range slices{
			s = append(s, val...)
	}
	return s
}

func TestConcat(t *testing.T) {
	test := []struct {
		slices   [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	}
	for _, tc := range test {
		assert.Equal(t, tc.expected, Concat(tc.slices))
	}
}
