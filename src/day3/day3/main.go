package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := readTestData("src/day3/day3/test-data.txt")
	if err != nil {
		fmt.Println("Error reading test data: ", err)
		return
	}
	fmt.Println("Total: ", SumEnabledMul(data))
}

func SumEnabledMul(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)
	total := 0
	enabled := true
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			total += num1 * num2
		}
	}
	return total
}

func SumMul(input string) int {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	total := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total += num1 * num2
	}
	return total
}

func readTestData(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
