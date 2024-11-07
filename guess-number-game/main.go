package main

import (
	"fmt"
	"math/rand"
)

// The main function is the entry point of the program
func main() {
	// Call the play function to start the game
	play()
}

// play function generates a random number and allows the user to guess it
func play() {
	// Generate a random number between 0 and 99
	randomNumber := rand.Intn(100)

	var userInput int
	var attempts int
	const maxAttempts = 10

	// Loop until the user reaches the maximum number of attempts
	for attempts < maxAttempts {
		attempts++
		// Prompt the user to input a number and display how many attempts are left
		fmt.Printf("Enter a number (remaining attempts: %d): ", maxAttempts-attempts+1)
		fmt.Scanln(&userInput)

		// Check if the user's input is correct
		if userInput == randomNumber {
			// If correct, congratulate the user and ask if they want to play again
			fmt.Println("Congratulations! You guessed the number.")
			askToPlayAgain()
			return
		} else if userInput < randomNumber {
			// If the guess is too low, inform the user
			fmt.Println("The number to guess is higher.")
		} else if userInput > randomNumber {
			// If the guess is too high, inform the user
			fmt.Println("The number to guess is lower.")
		}
	}

	// If the user runs out of attempts, reveal the correct number
	fmt.Println("You've run out of attempts. The number was:", randomNumber)
	askToPlayAgain()
}

// askToPlayAgain prompts the user whether they want to play again
func askToPlayAgain() {
	var option string

	// Ask the user if they want to play again
	fmt.Println("Do you want to play again? (y/n): ")
	fmt.Scanln(&option)

	// Handle the user's input
	switch option {
	case "y":
		// If yes, start a new game
		play()
	case "n":
		// If no, thank the user and exit
		fmt.Println("Thanks for playing!")
	default:
		// If the input is invalid, prompt again
		fmt.Println("Invalid option, please try again.")
		askToPlayAgain()
	}
}

// hello function returns a greeting message with the provided name
func hello(name string) string {
	return "Hello, " + name
}

// calc function returns the sum and multiplication of two integers
func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b

	return sum, mul
}

// calc1 function uses named return variables for sum and multiplication
func calc1(a, b int) (sum, mul int) {
	// Calculate sum and multiplication
	sum = a + b
	mul = a * b

	// Return the calculated values
	return
}
