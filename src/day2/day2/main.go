package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readTestData("src/day2/day2/test-data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Number of safe reports: ", CountSafeReports(reports))
}

func CountSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		output := checkIfReportIsSafe(report)
		if output {
			count++
		}
	}
	return count
}

func checkIfReportIsSafe(report []int) bool {
	output, i := checkReport(report)

	if !output {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])
		output, _ = checkReport(newReport)
	}

	if !output {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i-1])
		copy(newReport[i-1:], report[i:])
		output, _ = checkReport(newReport)
	}

	if !output {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[1:])
		output, _ = checkReport(newReport)
	}

	return output
}

func checkReport(report []int) (bool, int) {
	var increasing bool

	for i := 0; i < len(report)-1; i++ {
		signedDiff := report[i] - report[i+1]
		diff := int(math.Abs(float64(signedDiff)))

		if signedDiff == 0 || diff > 3 {
			return false, i + 1
		}

		if i == 0 {
			increasing = signedDiff > 0
		} else if increasing != (signedDiff > 0) {
			return false, i + 1
		}
	}
	return true, 0
}

func readTestData(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strValues := strings.Fields(line)
		var intValues []int
		for _, strValue := range strValues {
			intValue, err := strconv.Atoi(strValue)
			if err != nil {
				return nil, err
			}
			intValues = append(intValues, intValue)
		}
		data = append(data, intValues)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
