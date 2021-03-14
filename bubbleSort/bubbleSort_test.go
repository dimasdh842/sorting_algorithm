package bubbleSort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	result := Sort([]int{5, 2, 4, 3, 1, 0}, true)

	if !reflect.DeepEqual(result, []int{5, 4, 3, 2, 1, 0}) {
		t.Errorf("%v must be the same with [5, 4, 3, 2, 1, 0] ", result)
	}

	result = Sort([]int{5, 2, 4, 3, 1, 0}, false)

	if !reflect.DeepEqual(result, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("%v must be the same with [0, 1, 2, 3, 4, 5]", result)
	}
}
