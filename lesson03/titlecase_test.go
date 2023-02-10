package main

import (
	"testing"

	"github.com/kulti/titlecase"
)

func TestEmpty(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}
}
func TestWithoutMinor(t *testing.T) {
	const str, minor, want = "My baby a lost", "a lost", "My Baby a lost"
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}
}
func TestWithMinorInFirst(t *testing.T) {
	const str, minor, want = "my baby a lost", "my baby", "My baby A Lost"
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}

}
func TestWithNumber(t *testing.T) {
	const str, minor, want = "My baby a lost 4 toys", "a lost", "My Baby a lost 4 Toys"
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}

}
func TestWithChair(t *testing.T) {
	const str, minor, want = "My baby a lost _ toys", "_ toys", "My Baby A Lost _ toys"
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}

}
