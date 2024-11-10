package unittest

// Sum takes two integers a and b, and returns their sum.
func Sum(a, b int) int {
	return a + b
}

// GetMax takes two integers a and b, and returns the greater of the two.
func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Fibonacci calculates the nth Fibonacci number recursively.
// The function returns n if n is 0 or 1, and for other values,
// it recursively adds the two preceding numbers in the sequence.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
