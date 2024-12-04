package main

import (
	"strings"
	"testing"
)

func TestCheckForXMAS(t *testing.T) {
	expected := 18
	input :=
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	actual := CheckForXMAS(strings.Split(input, "\n"))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCheckForMasInX(t *testing.T) {
	expected := 9
	input :=
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	actual := CheckForMasInX(strings.Split(input, "\n"))
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
