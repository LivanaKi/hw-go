package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInt(t *testing.T){
	tests := []struct{
		str string
		expected int64
	}{
		{"-128", -128},
		{"0", 0},
		{"127", 127},
	}
	for _, tc := range tests{
		got, err := strconv.ParseInt(tc.str, 10, 8)
		require.NoError(t, err)
		require.EqualValues(t, tc.expected, got)
	}
}
func TestParseIntErrors(t *testing.T){
	tests := []string{"129", "128", "byaka"}
	for _, str  := range tests{
		_, err := strconv.ParseInt(str, 10, 8)
		require.Error(t, err)
	}
}