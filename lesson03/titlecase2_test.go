package main

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/require"
)

func TestEmpty1(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	require.Equal(t, want, got)
}
func TestWithoutMinor1(t *testing.T) {
	const str, minor, want = "My baby a lost", "a lost", "My Baby a lost"
	got := titlecase.TitleCase(str, minor)
	require.Equal(t, want, got)
}
func TestWithMinorInFirst1(t *testing.T) {
	const str, minor, want = "my baby a lost", "my baby", "My baby A Lost"
	got := titlecase.TitleCase(str, minor)
	require.Equal(t, want, got)

}
func TestWithNumber1(t *testing.T) {
	const str, minor, want = "My baby a lost 4 toys", "a lost", "My Baby a lost 4 Toys"
	got := titlecase.TitleCase(str, minor)
	require.Equal(t, want, got)

}
func TestWithChair1(t *testing.T) {
	const str, minor, want = "My baby a lost _ toys", "_ toys", "My Baby A Lost _ toys"
	got := titlecase.TitleCase(str, minor)
	require.Equal(t, want, got)
}
