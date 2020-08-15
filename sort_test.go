package main

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	unsorted := []int{2, 4, 3, 6, 1}
	fmt.Println(unsorted)

	sorted := mergeSort(unsorted)
	fmt.Println(sorted)

	// quickSort(unsorted, 0, len(unsorted)-1)
	// fmt.Println(unsorted)

	// sort.Ints(unsorted)
	// fmt.Println(unsorted)
}
