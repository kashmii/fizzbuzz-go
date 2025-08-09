package fizzbuzz

import "strconv"

// Convert converts an integer to its fizzbuzz representation.
// Returns "fizzbuzz" for multiples of 15, "fizz" for multiples of 3,
// "buzz" for multiples of 5, and the number as string otherwise.
func Convert(n int) string {
	if n%15 == 0 {
		return "fizzbuzz"
	}
	if n%3 == 0 {
		return "fizz"
	}
	if n%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(n)
}