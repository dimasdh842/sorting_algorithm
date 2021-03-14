package bubbleSort

func sweep(numbers []int, prevPhases int, reversed bool) bool {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool

	for secondIndex < (N - prevPhases) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if reversed {
			if firstNumber < secondNumber {
				numbers[firstIndex] = secondNumber
				numbers[secondIndex] = firstNumber
				didSwap = true
			}

		} else {
			if firstNumber > secondNumber {
				numbers[firstIndex] = secondNumber
				numbers[secondIndex] = firstNumber
				didSwap = true
			}

		}

		firstIndex++
		secondIndex++
	}

	return didSwap
}

func Sort(numbers []int, reversed bool) []int {
	N := len(numbers)
	i := 0
	for i < N {
		sweep(numbers, i, reversed)
		i++
	}

	return numbers
}
