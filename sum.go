package main

// Sum calculates the total from an array of numbers.
func Sum(numbers [5]int) int {
	sum := 0
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}
