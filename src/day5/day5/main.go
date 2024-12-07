package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, updates, err := readTestData("src/day5/day5/test-data.txt")
	if err != nil {
		fmt.Println("Error reading test data:", err)
		return
	}

	correctTotal, incorrectTotal := ProcessUpdates(rules, updates)

	fmt.Println("Correct total: ", correctTotal)
	fmt.Println("Incorrect total: ", incorrectTotal)
}

func ProcessUpdates(rules map[int][]int, updates [][]int) (int, int) {
	correctTotal := 0
	incorrectTotal := 0
	incorrectUpdates := [][]int{}

	for _, update := range updates {
		out := ProcessCorrectUpdate(rules, update)
		correctTotal += out
		if out == 0 {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	for _, update := range incorrectUpdates {
		incorrectTotal += ProcessIncorrectUpdate(rules, update)
	}

	return correctTotal, incorrectTotal
}

func ProcessIncorrectUpdate(rules map[int][]int, update []int) int {
	unavailablePages := make(map[int]struct{})
	sortedPages := []int{}

	for _, page := range update {
		_, exists := unavailablePages[page]
		if exists {
			insertPos := 0
			for i, p := range sortedPages {
				if contains(rules[page], p) {
					insertPos = i
					fmt.Println("insertPos: ", insertPos)
					fmt.Println("p: ", p)
					break
				}
			}

			newSorted := make([]int, len(sortedPages)+1)

			copy(newSorted, sortedPages[:insertPos])
			newSorted[insertPos] = page
			copy(newSorted[insertPos+1:], sortedPages[insertPos:])

			sortedPages = newSorted
		} else {
			sortedPages = append(sortedPages, page)
		}

		for _, rule := range rules[page] {
			unavailablePages[rule] = struct{}{}
		}
	}
	fmt.Println("sortedPages: ", sortedPages)

	return sortedPages[len(sortedPages)/2]
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func sortUpdate(rules map[int][]int, update []int) []int {
	unavailablePages := make(map[int]struct{})
	sortedPages := []int{}

	for _, page := range update {
		_, exists := unavailablePages[page]
		if exists {
			sortedPages = append([]int{page}, sortedPages...)
		} else {
			sortedPages = append(sortedPages, page)
		}

		for _, rule := range rules[page] {
			unavailablePages[rule] = struct{}{}
		}
	}

	return sortedPages
}

func ProcessCorrectUpdate(rules map[int][]int, update []int) int {
	unavailablePages := make(map[int]struct{})

	for _, page := range update {
		_, exists := unavailablePages[page]
		if exists {
			return 0
		}

		for _, rule := range rules[page] {
			unavailablePages[rule] = struct{}{}
		}
	}

	return update[len(update)/2]
}

func readTestData(filename string) (map[int][]int, [][]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	rules, updates, err := ProcessTestData(string(data))
	if err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}

func ProcessTestData(input string) (map[int][]int, [][]int, error) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(sections) != 2 {
		return nil, nil, fmt.Errorf("expected 2 sections, got %d", len(sections))
	}

	rules := map[int][]int{}
	for _, line := range strings.Split(sections[0], "\n") {
		numbers := strings.Split(line, "|")
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("invalid pair format: %s", line)
		}

		left, err := strconv.Atoi(strings.TrimSpace(numbers[0]))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid left number: %s", numbers[0])
		}

		right, err := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if err != nil {
			return nil, nil, fmt.Errorf("invalid right number: %s", numbers[1])
		}

		rules[right] = append(rules[right], left)
	}

	updates := [][]int{}
	for _, line := range strings.Split(sections[1], "\n") {
		var numbers []int
		for _, numStr := range strings.Split(line, ",") {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, nil, fmt.Errorf("invalid number in list: %s", numStr)
			}
			numbers = append(numbers, num)
		}
		updates = append(updates, numbers)
	}

	return rules, updates, nil
}
