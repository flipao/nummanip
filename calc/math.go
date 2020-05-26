package calc

import "errors"

// Add returns the sum of two numbers
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("provide more than 2 arguments"), sum
	}

	for _, num := range numbers {
		sum = sum + num
	}

	return nil, sum
}
