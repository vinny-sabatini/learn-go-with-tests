package arrays

// Sum calculates the total from a slice of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numberSet := range numbersToSum {
		if len(numberSet) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numberSet[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
