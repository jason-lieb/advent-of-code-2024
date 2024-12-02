package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	list1, list2, err := readTestData("src/day1/day1/test-data.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	heap1 := IntHeap(list1)
	heap2 := IntHeap(list2)
	heap.Init(&heap1)
	heap.Init(&heap2)
	fmt.Println("Distance:", FindDistance(&heap1, &heap2))

	instanceMap := CreateInstanceMap(list1, list2)

	fmt.Println("Similarity:", FindSimilarity(list1, instanceMap))
}

func FindDistance(heap1, heap2 *IntHeap) int {
	totalDistance := 0
	for heap1.Len() > 0 && heap2.Len() > 0 {
		val1 := heap.Pop(heap1).(int)
		val2 := heap.Pop(heap2).(int)
		distance := int(math.Abs(float64(val1 - val2)))
		totalDistance += distance
	}
	return totalDistance
}

func CreateInstanceMap(list1, list2 []int) map[int]int {
	set := make(map[int]struct{})
	for _, val := range list1 {
		set[val] = struct{}{}
	}
	instanceMap := make(map[int]int)
	for val := range set {
		instanceMap[val] = FindNumberOfInstances(val, list2)
	}
	return instanceMap
}

func FindNumberOfInstances(x int, list []int) int {
	count := 0
	for _, val := range list {
		if val == x {
			count++
		}
	}
	return count
}

func FindSimilarity(list []int, instanceMap map[int]int) int {
	similarity := 0
	for _, val := range list {
		similarity += instanceMap[val] * val
	}
	return similarity
}

func readTestData(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	return list1, list2, scanner.Err()
}
