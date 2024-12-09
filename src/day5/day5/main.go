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

func ProcessUpdates(rules [][]int, updates [][]int) (int, int) {
	correctTotal := 0
	incorrectTotal := 0
	incorrectUpdates := [][]int{}

	ruleMap := mkRuleMap(rules)
	for _, update := range updates {
		out := ProcessCorrectUpdate(ruleMap, update)
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

func ProcessIncorrectUpdate(rules [][]int, update []int) int {
	sortedPages := []int{}
	filteredRules := [][]int{}
	for _, rule := range rules {
		if contains(update, rule[0]) && contains(update, rule[1]) {
			filteredRules = append(filteredRules, rule)
		}
	}

	adjList := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, page := range update {
		adjList[page] = []int{}
		inDegree[page] = 0
	}
	for _, rule := range filteredRules {
		u, v := rule[0], rule[1]
		adjList[u] = append(adjList[u], v)
		inDegree[v]++
	}

	queue := []int{}
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		sortedPages = append(sortedPages, current)

		for _, neighbor := range adjList[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

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

func readTestData(filename string) ([][]int, [][]int, error) {
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

func mkRuleMap(rules [][]int) map[int][]int {
	ruleMap := map[int][]int{}
	for _, rule := range rules {
		u, v := rule[0], rule[1]
		ruleMap[v] = append(ruleMap[v], u)
	}
	return ruleMap
}

func ProcessTestData(input string) ([][]int, [][]int, error) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(sections) != 2 {
		return nil, nil, fmt.Errorf("expected 2 sections, got %d", len(sections))
	}

	rules := [][]int{}
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

		rules = append(rules, []int{left, right})
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
