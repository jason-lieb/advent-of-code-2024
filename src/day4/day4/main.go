package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rows, err := readTestData("src/day4/day4/test-data.txt")
	if err != nil {
		fmt.Println("Error reading test data: ", err)
		return
	}

	xmasTotal := CheckForXMAS(rows)
	masInXTotal := CheckForMasInX(rows)
	fmt.Println("XMAS Total: ", xmasTotal)
	fmt.Println("X-MAS Total: ", masInXTotal)
}

func CheckForMasInX(rows []string) int {
	total := 0
	for y, row := range rows {
		for x, cell := range row {
			if cell == 'A' {
				total += checkA(x, y, rows)
			}
		}
	}
	return total
}

func checkA(x int, y int, rows []string) int {
	if y-1 < 0 || y+1 >= len(rows) || x-1 < 0 || x+1 >= len(rows[0]) {
		return 0
	}

	switch {
	case rows[y+1][x+1] == 'M' && rows[y-1][x+1] == 'M' && rows[y+1][x-1] == 'S' && rows[y-1][x-1] == 'S':
		return 1
	case rows[y+1][x+1] == 'S' && rows[y-1][x+1] == 'S' && rows[y+1][x-1] == 'M' && rows[y-1][x-1] == 'M':
		return 1
	case rows[y+1][x+1] == 'S' && rows[y-1][x+1] == 'M' && rows[y+1][x-1] == 'S' && rows[y-1][x-1] == 'M':
		return 1
	case rows[y+1][x+1] == 'M' && rows[y-1][x+1] == 'S' && rows[y+1][x-1] == 'M' && rows[y-1][x-1] == 'S':
		return 1
	default:
		return 0
	}
}

func CheckForXMAS(rows []string) int {
	total := 0
	for y, row := range rows {
		for x, cell := range row {
			if cell == 'X' {
				total += checkX(x, y, rows)
			}
		}
	}
	return total
}

func checkX(x int, y int, rows []string) int {
	vertical := checkDirectionForX(x, y, 0, 1, rows) + checkDirectionForX(x, y, 0, -1, rows)
	horizontal := checkDirectionForX(x, y, 1, 0, rows) + checkDirectionForX(x, y, -1, 0, rows)
	diagonalLeft := checkDirectionForX(x, y, -1, 1, rows) + checkDirectionForX(x, y, -1, -1, rows)
	diagonalRight := checkDirectionForX(x, y, 1, 1, rows) + checkDirectionForX(x, y, 1, -1, rows)
	return vertical + horizontal + diagonalLeft + diagonalRight
}

func checkDirectionForX(x int, y int, dx int, dy int, rows []string) int {
	if y+3*dy < 0 || y+3*dy >= len(rows) || x+3*dx < 0 || x+3*dx >= len(rows[0]) {
		return 0
	}
	if rows[y+dy][x+dx] != 'M' {
		return 0
	}
	if rows[y+2*dy][x+2*dx] != 'A' {
		return 0
	}
	if rows[y+3*dy][x+3*dx] != 'S' {
		return 0
	}
	return 1
}

func readTestData(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	return rows, nil
}
