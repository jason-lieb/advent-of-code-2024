package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	total int
	args  []int
}

type Operation int

const (
	ADD = iota
	MULTIPLY
	CONCAT
)

func (o Operation) Compute(a, b int) int {
	switch o {
	case ADD:
		return a + b
	case MULTIPLY:
		return a * b
	case CONCAT:
		stringA := strconv.Itoa(a)
		stringB := strconv.Itoa(b)
		num, err := strconv.Atoi(stringA + stringB)
		if err != nil {
			panic(err)
		}
		return num
	default:
		panic(fmt.Sprintf("unknown operation: %v", o))
	}
}

func main() {
	data, err := os.ReadFile("src/day7/day7/test-data.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	rows := ProcessInput(string(data))

	fmt.Println("Total with ADD, MULTIPLY: ", CheckRows(rows, []Operation{ADD, MULTIPLY}))
	fmt.Println("Total with ADD, MULTIPLY, CONCAT: ", CheckRows(rows, []Operation{ADD, MULTIPLY, CONCAT}))
}

func CheckRows(rows []Row, operators []Operation) int {
	total := 0
	for _, row := range rows {
		total += checkRow(row, operators)
	}
	return total
}

func checkRow(row Row, operators []Operation) int {
	perms := generateOperationPermutations(len(row.args)-1, operators)

	for _, perm := range perms {
		total := row.args[0]
		for i, op := range perm {
			total = op.Compute(total, row.args[i+1])
		}
		if total == row.total {
			return total
		}
	}

	return 0
}

func generateOperationPermutations(length int, operators []Operation) [][]Operation {
	result := make([][]Operation, 0)

	var generate func(current []Operation)
	generate = func(current []Operation) {
		if len(current) == length {
			temp := make([]Operation, length)
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for _, op := range operators {
			generate(append(current, op))
		}
	}

	generate([]Operation{})
	return result
}

func ProcessInput(data string) []Row {
	rows := []Row{}
	lines := strings.Split(data, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return rows
		}
		argsArray := strings.Split(strings.TrimSpace(parts[1]), " ")

		total, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		args := []int{}
		for _, arg := range argsArray {
			num, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			args = append(args, num)
		}

		rows = append(rows, Row{
			total: total,
			args:  args,
		})
	}

	return rows
}
