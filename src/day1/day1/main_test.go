package main

import (
	"container/heap"
	"testing"
)

func TestFindDistance(t *testing.T) {
	expected := 11

	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	heap1 := IntHeap(list1)
	heap2 := IntHeap(list2)
	heap.Init(&heap1)
	heap.Init(&heap2)

	result := FindDistance(&heap1, &heap2)

	if result != expected {
		t.Errorf("MyFunction() = %v; want %v", result, expected)
	}
}

func TestFindSimilarity(t *testing.T) {
	expected := 31

	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	instanceMap := CreateInstanceMap(list1, list2)

	result := FindSimilarity(list1, instanceMap)

	if result != expected {
		t.Errorf("MyFunction() = %v; want %v", result, expected)
	}
}
