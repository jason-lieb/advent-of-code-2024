package main

import (
	"fmt"
	"os"
	"strings"
)

type Location struct {
	X int
	Y int
}

func (l Location) Add(other Location) Location {
	return Location{
		X: l.X + other.X,
		Y: l.Y + other.Y,
	}
}

func (l Location) Subtract(other Location) Location {
	return Location{
		X: l.X - other.X,
		Y: l.Y - other.Y,
	}
}

func main() {
	data, err := os.ReadFile("src/day8/day8/test-data.txt")
	if err != nil {
		panic(err)
	}

	grid := ProcessInput(string(data))
	bounds := Location{X: len(grid[0]), Y: len(grid)}
	antinodeLocations := make(map[Location]struct{})
	resonantAntinodeLocations := make(map[Location]struct{})
	antennas := FindAntennas(grid)

	total := 0
	resonantTotal := 0
	for _, antenna := range antennas {
		total += FindNumberOfAntinodes(antenna, bounds, antinodeLocations)
		resonantTotal += FindNumberOfResonantAntinodes(antenna, bounds, resonantAntinodeLocations)
	}

	fmt.Println("Number of Antinodes:", total)
	fmt.Println("Number of Resonant Antinodes:", resonantTotal)
}

func FindNumberOfResonantAntinodes(antennas []Location, bounds Location, antinodeLocations map[Location]struct{}) int {
	antinodesCount := 0

	if len(antennas) == 1 {
		return 0
	}

	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			diff := antennas[i].Subtract(antennas[j])
			if _, exists := antinodeLocations[antennas[i]]; !exists {
				antinodesCount++
				antinodeLocations[antennas[i]] = struct{}{}
			}
			if _, exists := antinodeLocations[antennas[j]]; !exists {
				antinodesCount++
				antinodeLocations[antennas[j]] = struct{}{}
			}
			antinodesCount += FindResonantAntinodes(antennas[i], diff, Location.Add, bounds, antinodeLocations)
			antinodesCount += FindResonantAntinodes(antennas[j], diff, Location.Subtract, bounds, antinodeLocations)
		}
	}

	return antinodesCount
}

func FindResonantAntinodes(antenna Location, diff Location, operation func(Location, Location) Location, bounds Location, antinodeLocations map[Location]struct{}) int {
	antinodesCount := 0
	currentAntinode := antenna

	for {
		nextAntinode := operation(currentAntinode, diff)

		if nextAntinode.X < 0 || nextAntinode.X >= bounds.X ||
			nextAntinode.Y < 0 || nextAntinode.Y >= bounds.Y {
			break
		}

		if _, exists := antinodeLocations[nextAntinode]; !exists {
			antinodesCount++
			antinodeLocations[nextAntinode] = struct{}{}
		}

		currentAntinode = nextAntinode
	}

	return antinodesCount
}

func FindNumberOfAntinodes(antennas []Location, bounds Location, antinodeLocations map[Location]struct{}) int {
	antinodesCount := 0

	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			diff := antennas[i].Subtract(antennas[j])
			addAntinode := antennas[i].Add(diff)
			subtractAntinode := antennas[j].Subtract(diff)
			if addAntinode.X >= 0 && addAntinode.X < bounds.X && addAntinode.Y >= 0 && addAntinode.Y < bounds.Y {
				if _, exists := antinodeLocations[addAntinode]; !exists {
					antinodesCount++
					antinodeLocations[addAntinode] = struct{}{}
				}
			}
			if subtractAntinode.X >= 0 && subtractAntinode.X < bounds.X && subtractAntinode.Y >= 0 && subtractAntinode.Y < bounds.Y {
				if _, exists := antinodeLocations[subtractAntinode]; !exists {
					antinodesCount++
					antinodeLocations[subtractAntinode] = struct{}{}
				}
			}
		}
	}

	return antinodesCount
}

func FindAntennas(grid [][]rune) map[rune][]Location {
	antennas := make(map[rune][]Location)

	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], Location{X: x, Y: y})
			}
		}
	}

	return antennas
}

func ProcessInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	return grid
}
