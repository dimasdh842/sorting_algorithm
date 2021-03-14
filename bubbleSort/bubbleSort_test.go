package bubbleSort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	numbers := []int{5, 1, 4, 3, 2, 0}
	result := Sort(numbers, true)

	if reflect.DeepEqual(result, []int{5, 4, 3, 2, 1, 0}) {
		t.Error("Result must be the same ")
	}

	result = Sort(numbers, false)

	if reflect.DeepEqual(result, []int{0, 1, 2, 3, 4, 5}) {
		t.Error("Result must be the same ")
	}

}
