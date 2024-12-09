package main

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	input :=
		`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`
	expectedTarget := Position{x: 4, y: 6}
	expectedDirection := Up
	actualTarget, actualDirection, _ := ParseInput(input)

	if actualTarget != expectedTarget {
		t.Errorf("Expected Target: %v, got %v", expectedTarget, actualTarget)
	}
	if actualDirection != expectedDirection {
		t.Errorf("Expected Direction: %v, got %v", expectedDirection, actualDirection)
	}
}

func TestCountVisitedPositions(t *testing.T) {
	input :=
		`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	target, direction, grid := ParseInput(input)
	actual, _ := CountVisitedPositions(target, direction, grid)
	expected := 41

	if actual != expected {
		t.Errorf("Expected Visited Positions: %d, got %d", expected, actual)
	}

	expectedObstacleCount := 6

	obstacleLocations := FindPossibleObstacleLocations(grid)
	count := 0
	for _, obstacleLocation := range obstacleLocations {
		newGrid := make(Grid, len(grid))
		for i := range grid {
			newGrid[i] = make([]rune, len(grid[i]))
			copy(newGrid[i], grid[i])
		}
		newGrid[obstacleLocation.y][obstacleLocation.x] = '#'
		_, loop := CountVisitedPositions(target, direction, newGrid)
		if loop {
			count++
		}
	}

	if count != expectedObstacleCount {
		t.Errorf("Expected Obstacle Count: %d, got %d", expectedObstacleCount, count)
	}
}
