package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// PrintList accepts a variable number of arguments (of any type)
// and prints each item in the list.
func PrintList(list ...any) {
	// Iterate through each element in the list
	for _, value := range list {
		// Print each value in the list
		fmt.Println(value)
	}
}

// Define a custom type `integer` which is an alias for `int`.
type integer int

// Define an interface `Numbers` with constraints.
// This interface allows types that are either int, float64, float32, or uint.
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// Includes checks if a value exists in a list of comparable type `T`.
// It returns true if the value is found, otherwise false.
func Includes[T comparable](list []T, value T) bool {
	// Iterate over the list of elements
	for _, item := range list {
		// If a match is found, return true
		if item == value {
			return true
		}
	}
	// Return false if no match is found
	return false
}

// Sum accepts a variable number of numeric arguments of a constrained type `T` (either Integer or Float),
// and returns their sum.
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	// Iterate through the numbers and add them to the total
	for _, num := range nums {
		total += num
	}
	// Return the final total
	return total
}

// Filter filters a list of ordered elements (such as int, string, etc.)
// based on the provided callback function. It returns a new list with the elements
// that satisfy the condition in the callback function.
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	// Create an empty slice to hold the filtered result
	result := make([]T, 0, len(list))
	// Iterate through each item in the list
	for _, item := range list {
		// Apply the callback function to filter the item
		if callback(item) {
			// If the callback returns true, append the item to the result
			result = append(result, item)
		}
	}
	// Return the filtered list
	return result
}

// Define a generic struct `Product` that can hold either uint or string as the type for the product `Id`.
// This allows flexibility in the type of the product identifier.
type Product[T uint | string] struct {
	Id    T       // Product identifier (can be uint or string)
	Desc  string  // Description of the product
	Price float32 // Price of the product
}

func main() {
	// Example usage of PrintList (commented out for now)
	// PrintList("John", 34, 5.5, true)

	// Example usage of custom integer type `integer`
	// var num1 integer = 100
	// var num2 integer = 300
	// fmt.Println(Sum(4, 8, 9, 1.5)) // Sum with mixed types (int and float)
	// fmt.Println(Sum(num1, num2)) // Sum with custom integer type

	// Example usage of Includes with a list of strings and numbers
	// strings := []string{"a", "b", "c", "d"}
	// numbers := []int{1, 2, 3, 4}

	// fmt.Println(Includes(strings, "a")) // true
	// fmt.Println(Includes(strings, "f")) // false
	// fmt.Println(Includes(numbers, 4)) // true
	// fmt.Println(Includes(numbers, 8)) // false

	// Create and print two products with different types for the product Id
	product1 := Product[uint]{1, "shoes", 50}
	product2 := Product[string]{"FD-ASDF", "shoes", 50}
	fmt.Println(product1, product2)

	// Example usage of Filter function (commented out for now)
	// fmt.Println(Filter(numbers, func(value int) bool { return value > 3 })) // Filter values greater than 3
	// fmt.Println(Filter(strings, func(value string) bool { return value > "b" })) // Filter strings lexicographically greater than "b"
}
