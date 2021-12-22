package main

func SumAll(numbersToSum ...[]int) []int {
	lenOfNumbers := len(numbersToSum)
	sums := make([]int, lenOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
