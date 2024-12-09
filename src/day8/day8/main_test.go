package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	input :=
		`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	expectedNumberOfAntinodes := 14
	expectedNumberOfResonantAntinodes := 34
	grid := ProcessInput(input)
	bounds := Location{X: len(grid[0]), Y: len(grid)}
	antinodeLocations := make(map[Location]struct{})
	resonantAntinodeLocations := make(map[Location]struct{})

	antennas := FindAntennas(grid)

	actualNumberOfAntinodes := 0
	actualNumberOfResonantAntinodes := 0
	for _, antenna := range antennas {
		actualNumberOfAntinodes += FindNumberOfAntinodes(antenna, bounds, antinodeLocations)
		actualNumberOfResonantAntinodes += FindNumberOfResonantAntinodes(antenna, bounds, resonantAntinodeLocations)
	}

	if actualNumberOfAntinodes != expectedNumberOfAntinodes {
		t.Errorf("Expected %d, got %d", expectedNumberOfAntinodes, actualNumberOfAntinodes)
	}
	if actualNumberOfResonantAntinodes != expectedNumberOfResonantAntinodes {
		t.Errorf("Expected %d, got %d", expectedNumberOfResonantAntinodes, actualNumberOfResonantAntinodes)
	}
}
