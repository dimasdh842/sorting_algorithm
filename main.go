package main

import (
	"fmt"

	"github.com/dimasdh842/sorting_algorithm/bubbleSort"
)

func main() {

	numbers := []int{5, 1, 4, 3, 2, 0}
	fmt.Println("before swap : ", numbers)
	bubbleSort.Sort(numbers)
	fmt.Println("after sorting : ", numbers)

}
