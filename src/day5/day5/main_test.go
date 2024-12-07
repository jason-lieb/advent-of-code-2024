package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	input := `
	47|53
	97|13
	97|61
	97|47
	75|29
	61|13
	75|53
	29|13
	97|29
	53|29
	61|53
	97|53
	61|29
	47|13
	75|47
	97|75
	47|61
	75|61
	47|29
	75|13
	53|13

	75,47,61,53,29
	97,61,53,29,13
	75,29,13
	75,97,47,61,53
	61,13,29
	97,13,75,29,47`

	expectedCorrectTotal := 143
	expectedIncorrectTotal := 123

	rules, updates, err := ProcessTestData(input)
	if err != nil {
		t.Errorf("Error processing test data: %s", err)
	}

	actualCorrectTotal, actualIncorrectTotal := ProcessUpdates(rules, updates)

	if actualCorrectTotal != expectedCorrectTotal {
		t.Errorf("Expected Correct Total: %d, got %d", expectedCorrectTotal, actualCorrectTotal)
	}

	if actualIncorrectTotal != expectedIncorrectTotal {
		t.Errorf("Expected Incorrect Total: %d, got %d", expectedIncorrectTotal, actualIncorrectTotal)
	}
}
