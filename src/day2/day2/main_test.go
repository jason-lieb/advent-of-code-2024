package main

import "testing"

func TestCountSafeReports(t *testing.T) {
	expected := 4

	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := CountSafeReports(reports)

	if result != expected {
		t.Errorf("MyFunction() = %v; want %v", result, expected)
	}
}
