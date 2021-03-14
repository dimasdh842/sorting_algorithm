package main

import (
	"fmt"
	"reflect"

	"github.com/dimasdh842/sorting_algorithm/bubbleSort"
)

func main() {

	numbers := []int{5, 1, 4, 3, 2, 0}
	result := bubbleSort.Sort(numbers, true)
	fmt.Println(result)
	fmt.Println(numbers)

	// sliceA := []int{1, 2, 3, 11}
	// sliceB := []int{1, 2, 3, 4}

	if reflect.DeepEqual(result, numbers) {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}

}
