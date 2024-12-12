package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	file = iota
	freeSpace
)

type Chunk struct {
	Value         int
	Count         int
	TrailingSpace int
	Index         int
}

type Chunks []Chunk

func (chunks Chunks) TotalCount() int {
	total := 0
	for _, chunk := range chunks {
		total += chunk.Count
	}
	return total
}

func main() {
	data, err := os.ReadFile("src/day9/day9/test-data.txt")
	if err != nil {
		panic(err)
	}

	diskmap := string(data)
	diskmap = strings.Map(func(r rune) rune {
		if r == '\n' {
			return -1
		}
		return r
	}, diskmap)

	hardDrive := BuildHardDrive([]rune(diskmap))
	compressed := CompressData(hardDrive)
	compressedInChunks := CompressDataInChunks(hardDrive)
	checksum := CalculateChecksum(compressed)
	checksumInChunks := CalculateChecksum(compressedInChunks)

	fmt.Println("Checksum:", checksum)
	fmt.Println("Checksum in chunks:", checksumInChunks)
}

func CalculateChecksum(compressed []int) int {
	checksum := 0
	vals := []int{}
	for i := 0; i < len(compressed); i++ {
		if compressed[i] != -1 {
			val := compressed[i] * i
			vals = append(vals, val)
			checksum += val
		}
	}
	return checksum
}

func CompressDataInChunks(hardDrive []int) []int {
	hardDriveChunks := Chunks{}
	lastValue := 0
	lastCount := 0
	index := 0
	for i := 0; i < len(hardDrive); i++ {
		if hardDrive[i] != lastValue {
			hardDriveChunks = append(hardDriveChunks, Chunk{Value: lastValue, Count: lastCount, TrailingSpace: 0, Index: index})
			lastCount = 0
		}
		lastValue = hardDrive[i]
		lastCount++
		index++
	}
	hardDriveChunks = append(hardDriveChunks, Chunk{Value: lastValue, Count: lastCount, TrailingSpace: 0, Index: index})

	filteredReversedFiles := Chunks{}
	for i := len(hardDriveChunks) - 1; i > 0; i-- {
		if hardDriveChunks[i].Value != -1 {
			filteredReversedFiles = append(filteredReversedFiles, hardDriveChunks[i])
		}
	}

	for i := 0; i < len(filteredReversedFiles); i++ {
		for j := 0; j < len(hardDriveChunks); j++ {
			if hardDriveChunks[j].Value != -1 {
				continue
			}
			if hardDriveChunks[j].Index > filteredReversedFiles[i].Index {
				continue
			}
			if hardDriveChunks[j].Count == filteredReversedFiles[i].Count {
				removeMovedChunks(hardDriveChunks, filteredReversedFiles[i].Value)
				hardDriveChunks[j].Value = filteredReversedFiles[i].Value
				break
			} else if hardDriveChunks[j].Count > filteredReversedFiles[i].Count {
				removeMovedChunks(hardDriveChunks, filteredReversedFiles[i].Value)
				hardDriveChunks[j].Value = filteredReversedFiles[i].Value
				hardDriveChunks[j].TrailingSpace = hardDriveChunks[j].Count - filteredReversedFiles[i].Count
				hardDriveChunks[j].Count = filteredReversedFiles[i].Count
				break
			}
		}
		hardDriveChunks = flattenChunks(hardDriveChunks)
	}
	first10 := Chunks{}
	count := 0
	for i := 0; i < len(hardDriveChunks) && count < 10; i++ {
		first10 = append(first10, hardDriveChunks[i])
		count++
	}

	fmt.Println("Hard Drive Chunks", printHardDriveChunks(first10))
	return chunksToList(hardDriveChunks)
}

func printHardDriveChunks(hardDriveChunks Chunks) string {
	var result strings.Builder

	for _, chunk := range hardDriveChunks {
		for i := 0; i < chunk.Count; i++ {
			if chunk.Value == -1 {
				result.WriteString(".")
			} else {
				result.WriteString(strconv.Itoa(chunk.Value))
			}
		}
	}

	return result.String()
}

func flattenChunks(hardDriveChunks Chunks) Chunks {
	newHardDrive := Chunks{}
	for i := 0; i < len(hardDriveChunks); i++ {
		if hardDriveChunks[i].TrailingSpace > 0 {
			newChunks := append(Chunks{}, Chunk{Value: hardDriveChunks[i].Value, Count: hardDriveChunks[i].Count, TrailingSpace: 0})
			newChunks = append(newChunks, Chunk{Value: -1, Count: hardDriveChunks[i].TrailingSpace, TrailingSpace: 0})
			newHardDrive = append(newHardDrive, newChunks...)
		} else {
			newHardDrive = append(newHardDrive, hardDriveChunks[i])
		}
	}
	return newHardDrive
}

func removeMovedChunks(hardDriveChunks Chunks, value int) {
	for i := 0; i < len(hardDriveChunks); i++ {
		if hardDriveChunks[i].Value == value {
			hardDriveChunks[i].Value = -1
		}
	}
}

func chunksToList(hardDriveChunks Chunks) []int {
	list := []int{}
	for i := 0; i < len(hardDriveChunks); i++ {
		for j := 0; j < hardDriveChunks[i].Count; j++ {
			list = append(list, hardDriveChunks[i].Value)
		}
	}
	return list
}

func CompressData(hardDrive []int) []int {
	filteredReversedFiles := []int{}
	filteredFileIndex := 0

	for i := len(hardDrive) - 1; i >= 0; i-- {
		if hardDrive[i] != -1 {
			filteredReversedFiles = append(filteredReversedFiles, hardDrive[i])
		}
	}

	compressedFiles := []int{}
	for i := 0; i < len(filteredReversedFiles); i++ {
		if hardDrive[i] == -1 {
			compressedFiles = append(compressedFiles, filteredReversedFiles[filteredFileIndex])
			filteredFileIndex++
		} else {
			compressedFiles = append(compressedFiles, hardDrive[i])
		}
	}

	return compressedFiles
}

func BuildHardDrive(diskmap []rune) []int {
	hardDrive := []int{}
	id := 0
	space := file

	for i := 0; i < len(diskmap); i++ {
		val, err := strconv.Atoi(string(diskmap[i]))
		if err != nil {
			panic(err)
		}
		for j := 0; j < val; j++ {
			if space == file {
				hardDrive = append(hardDrive, id)
			} else {
				hardDrive = append(hardDrive, -1)
			}
		}
		if space == file {
			id++
			space = freeSpace
		} else {
			space = file
		}
	}

	return hardDrive
}
