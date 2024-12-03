package main

import "testing"

func TestSumMul(t *testing.T) {
	expected := 161
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := SumMul(input)
	if result != expected {
		t.Errorf("SumMul() = %v; want %v", result, expected)
	}
}

func TestSumEnabledMul(t *testing.T) {
	expected := 48
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := SumEnabledMul(input)
	if result != expected {
		t.Errorf("SumEnabledMul() = %v; want %v", result, expected)
	}
}
