package bubbleSort

func sweep(numbers []int, prevPhases int) bool {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool

	for secondIndex < (N - prevPhases) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}

	return didSwap
}

func Sort(numbers []int) []int {
	N := len(numbers)
	i := 0
	for i < N {
		if !sweep(numbers, i) {
			return []int{}
		}
		i++
	}

	return numbers
}
