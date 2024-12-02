package main

import "testing"

func TestCountSafeReports(t *testing.T) {
	expected := 6

	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
		{7, 10, 8, 10, 11},
		{29, 28, 27, 25, 26, 25, 22, 20},
	}

	result := CountSafeReports(reports)

	if result != expected {
		t.Errorf("MyFunction() = %v; want %v", result, expected)
	}
}
