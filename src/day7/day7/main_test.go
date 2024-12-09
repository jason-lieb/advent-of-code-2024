package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	input :=
		`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	rows := ProcessInput(input)

	expectedFromAddAndMultiply := 3749
	actualFromAddAndMultiply := CheckRows(rows, []Operation{ADD, MULTIPLY})
	if actualFromAddAndMultiply != expectedFromAddAndMultiply {
		t.Errorf("Expected %d, got %d", expectedFromAddAndMultiply, actualFromAddAndMultiply)
	}

	expectedFromAddAndMultiplyAndConcat := 11387
	actualFromAddAndMultiplyAndConcat := CheckRows(rows, []Operation{ADD, MULTIPLY, CONCAT})
	if actualFromAddAndMultiplyAndConcat != expectedFromAddAndMultiplyAndConcat {
		t.Errorf("Expected %d, got %d", expectedFromAddAndMultiplyAndConcat, actualFromAddAndMultiplyAndConcat)
	}
}
