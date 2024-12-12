package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	input := "2333133121414131402"
	expected := 1928
	expectedInChunks := 2858

	hardDrive := BuildHardDrive([]rune(input))
	compressed := CompressData(hardDrive)
	compressedInChunks := CompressDataInChunks(hardDrive)
	checksum := CalculateChecksum(compressed)
	checksumInChunks := CalculateChecksum(compressedInChunks)

	if checksum != expected {
		t.Errorf("Expected %d, got %d", expected, checksum)
	}

	if checksumInChunks != expectedInChunks {
		t.Errorf("Expected %d, got %d", expectedInChunks, checksumInChunks)
	}

	input2 := "12143"
	// expected2 := 1928
	expectedInChunks2 := 31

	hardDrive2 := BuildHardDrive([]rune(input2))
	compressedInChunks2 := CompressDataInChunks(hardDrive2)
	checksumInChunks2 := CalculateChecksum(compressedInChunks2)

	if checksumInChunks2 != expectedInChunks2 {
		t.Errorf("Expected %d, got %d", expectedInChunks2, checksumInChunks2)
	}
}
