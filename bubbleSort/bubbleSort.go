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
<<<<<<< HEAD
=======

>>>>>>> d697921f2206f787f491bee2a626760ae4f4b0c5
			if firstNumber < secondNumber {
				numbers[firstIndex] = secondNumber
				numbers[secondIndex] = firstNumber
				didSwap = true
			}
<<<<<<< HEAD

=======
>>>>>>> d697921f2206f787f491bee2a626760ae4f4b0c5
		} else {
			if firstNumber > secondNumber {
				numbers[firstIndex] = secondNumber
				numbers[secondIndex] = firstNumber
				didSwap = true
			}
<<<<<<< HEAD

=======
>>>>>>> d697921f2206f787f491bee2a626760ae4f4b0c5
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
<<<<<<< HEAD
		sweep(numbers, i, reversed)
=======
		if !sweep(numbers, i, reversed) {
			return []int{}
		}
>>>>>>> d697921f2206f787f491bee2a626760ae4f4b0c5
		i++
	}

	return numbers
}
