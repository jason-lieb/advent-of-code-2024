package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid [][]rune

func (g Grid) String() string {
	var result strings.Builder
	for _, row := range g {
		for _, cell := range row {
			result.WriteRune(cell)
		}
		result.WriteRune('\n')
	}
	return result.String()
}

type Position struct {
	x, y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[d]
}

func (d Direction) Move() Position {
	return [...]Position{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}[d]
}

func (d Direction) Turn() Direction {
	return Direction((int(d) + 1) % 4)
}

func main() {
	data, err := os.ReadFile("src/day6/day6/test-data.txt")
	if err != nil {
		panic(err)
	}

	currentPosition, direction, grid := ParseInput(string(data))
	total, _ := CountVisitedPositions(currentPosition, direction, grid)

	fmt.Println("Total Visited Positions:", total)

	obstacleLocations := FindPossibleObstacleLocations(grid)
	count := 0
	for _, obstacleLocation := range obstacleLocations {
		newGrid := make(Grid, len(grid))
		for i := range grid {
			newGrid[i] = make([]rune, len(grid[i]))
			copy(newGrid[i], grid[i])
		}
		newGrid[obstacleLocation.y][obstacleLocation.x] = '#'
		_, loop := CountVisitedPositions(currentPosition, direction, newGrid)
		if loop {
			count++
		}
	}

	fmt.Println("Number of possible obstacle locations: ", count)
}

func FindPossibleObstacleLocations(grid Grid) []Position {
	obstacleLocations := []Position{}
	for y, row := range grid {
		for x, cell := range row {
			if cell == '.' {
				obstacleLocations = append(obstacleLocations, Position{x, y})
			}
		}
	}
	return obstacleLocations
}

func CountVisitedPositions(currentPosition Position, direction Direction, grid Grid) (int, bool) {
	loop := false
	visited := map[Position]map[Direction]struct{}{}

	for {
		if _, exists := visited[currentPosition]; !exists {
			visited[currentPosition] = make(map[Direction]struct{})
		}
		if _, hasDirection := visited[currentPosition][direction]; hasDirection {
			loop = true
			break
		}
		visited[currentPosition][direction] = struct{}{}

		nextSpace := Position{currentPosition.x + direction.Move().x, currentPosition.y + direction.Move().y}

		inXBounds := nextSpace.x >= 0 && nextSpace.x < len(grid[0])
		inYBounds := nextSpace.y >= 0 && nextSpace.y < len(grid)
		if !inXBounds || !inYBounds {
			break
		}

		if grid[nextSpace.y][nextSpace.x] == '#' {
			direction = direction.Turn()
		} else {
			currentPosition = nextSpace
		}
	}

	return len(visited), loop
}

func ParseInput(input string) (Position, Direction, Grid) {
	target := Position{}
	var direction Direction

	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make(Grid, len(lines))

	for y, line := range lines {
		grid[y] = []rune(line)
		for x, char := range line {
			if char != '#' && char != '.' {
				target = Position{x, y}

				switch char {
				case '>':
					direction = Right
				case 'v':
					direction = Down
				case '<':
					direction = Left
				default:
					direction = Up
				}
			}
		}
	}
	return target, direction, grid
}
