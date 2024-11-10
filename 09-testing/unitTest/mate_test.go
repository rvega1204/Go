package unittest

import "testing"

// TestSum tests the Sum function with various pairs of integers and checks if their sum matches the expected result.
func TestSum(t *testing.T) {
	// Define a table of test cases with pairs of integers (a, b) and their expected sum (c).
	table := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	// Loop through each test case
	for _, item := range table {
		// Calculate the sum of a and b using the Sum function
		total := Sum(item.a, item.b)
		// If the result does not match the expected value, log an error
		if total != item.c {
			t.Errorf("Incorrect Sum, got %d, expected %d", total, item.c)
		}
	}
}

// TestGetMax tests the GetMax function to ensure it correctly identifies the maximum of two integers.
func TestGetMax(t *testing.T) {
	// Define a table of test cases with pairs of integers (a, b) and their expected maximum (c).
	table := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{5, 3, 5},
		{2, 3, 3},
	}

	// Loop through each test case
	for _, item := range table {
		// Get the maximum of a and b using the GetMax function
		max := GetMax(item.a, item.b)
		// If the result does not match the expected maximum, log an error
		if max != item.c {
			t.Errorf("Incorrect GetMax, got %d, expected %d", max, item.c)
		}
	}
}

// TestFibonacci tests the Fibonacci function by comparing its output for given inputs to expected results.
func TestFibonacci(t *testing.T) {
	// Define a table of test cases with the input number (n) and the expected Fibonacci result (r).
	table := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	// Loop through each test case
	for _, item := range table {
		// Calculate the Fibonacci number for n using the Fibonacci function
		fibo := Fibonacci(item.n)
		// If the result does not match the expected result, log an error
		if fibo != item.r {
			t.Errorf("Incorrect Fibonacci, got %d, expected %d", fibo, item.r)
		}
	}
}
