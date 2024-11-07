package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database. If an error occurs during the connection, log it and terminate the program.
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the database connection is closed when the program terminates.
	defer db.Close()

	// Infinite loop to display the menu and prompt the user for an action until they choose to exit.
	for {
		// Display the main menu to the user with options for different contact operations.
		fmt.Println("\nMenu:")
		fmt.Println("1. Contact list")       // Option 1: Show the list of contacts
		fmt.Println("2. Get contact by ID")  // Option 2: Retrieve a contact by its ID
		fmt.Println("3. Create new contact") // Option 3: Create a new contact
		fmt.Println("4. Update contact")     // Option 4: Update an existing contact
		fmt.Println("5. Delete contact")     // Option 5: Delete a contact
		fmt.Println("6. Exit")               // Option 6: Exit the program
		fmt.Println("Please select option: ")

		// Variable to store the user's option choice.
		var option int

		// Read the user's input for the selected option.
		// Scanln reads the option as an integer and assigns it to the variable 'option'.
		fmt.Scanln(&option)

		// Switch case to handle each option based on the user's choice.
		switch option {
		case 1:
			// Show the list of contacts when option 1 is selected.
			handlers.ListContacts(db)
		case 2:
			// Prompt the user for a contact ID to retrieve a specific contact by ID.
			fmt.Print("Enter the contact ID: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactByID(db, idContact)
		case 3:
			// Prompt the user to input details for a new contact (calls inputContactDetails to get data).
			// Then, create the new contact by calling the CreateContact handler function.
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			// After creating the contact, display the updated contact list.
			handlers.ListContacts(db)
		case 4:
			// Similar to option 3, but for updating an existing contact.
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			// After updating the contact, display the updated contact list.
			handlers.ListContacts(db)
		case 5:
			// Prompt the user for a contact ID to delete the contact.
			fmt.Print("Enter the contact ID to delete: ")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
			// After deleting the contact, display the updated contact list.
			handlers.ListContacts(db)
		case 6:
			// Option 6 is to exit the program.
			// Display a message and return, which terminates the loop and ends the program.
			fmt.Println("Leaving the program...")
			return
		default:
			// If the user enters an invalid option (not between 1 and 6), show an error message.
			fmt.Println("Invalid option, please select a valid option")
		}
	}
}

// Function to validate email format using a regular expression
func isValidEmail(email string) bool {
	// Simple regular expression to validate emails (it can be more complex if needed)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func inputContactDetails(option int) models.Contact {
	// Create a reader to read input from the user
	reader := bufio.NewReader(os.Stdin)
	var contact models.Contact

	// If the option is 4, ask for the contact ID
	if option == 4 {
		fmt.Print("Enter the contact ID: ")
		var idContact int

		// Try to read the ID and handle any input errors
		_, err := fmt.Scanln(&idContact)
		if err != nil {
			// If an error occurs, print a message and return
			fmt.Println("Invalid input. Please enter a valid number for the contact ID.")
			return models.Contact{} // Return an empty contact if the input is invalid
		}

		// Validate that the ID is a positive number, if necessary
		if idContact <= 0 {
			fmt.Println("Contact ID must be a positive number.")
			return models.Contact{} // Return an empty contact if the ID is invalid
		}

		// Assign the ID to the contact
		contact.Id = idContact
	}

	// Ask for the contact name, repeating if the input is invalid
	for {
		fmt.Print("Enter contact name: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading name:", err)
			continue // Continue asking for the name if there was an error
		}

		// Trim spaces from the beginning and end of the input
		name = strings.TrimSpace(name)

		// Check if the name is empty
		if name == "" {
			fmt.Println("Name cannot be empty. Please enter a valid name.")
			continue
		}

		// Assign the name to the contact and exit the loop
		contact.Name = name
		break
	}

	// Ask for the contact email, repeating if the input is invalid
	for {
		fmt.Print("Enter contact email: ")
		email, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading email:", err)
			continue // Continue asking for the email if there was an error
		}

		// Trim spaces from the beginning and end of the input
		email = strings.TrimSpace(email)

		// Check if the email is empty
		if email == "" {
			fmt.Println("Email cannot be empty. Please enter a valid email.")
			continue
		}

		// Validate that the email format is correct
		if !isValidEmail(email) {
			fmt.Println("Invalid email format. Please enter a valid email.")
			continue
		}

		// Assign the email to the contact and exit the loop
		contact.Email = email
		break
	}

	// Ask for the contact phone number, repeating if the input is invalid
	for {
		fmt.Print("Enter contact phone: ")
		phone, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading phone:", err)
			continue // Continue asking for the phone number if there was an error
		}

		// Trim spaces from the beginning and end of the input
		phone = strings.TrimSpace(phone)

		// Check if the phone number is empty
		if phone == "" {
			fmt.Println("Phone number cannot be empty. Please enter a valid phone number.")
			continue
		}

		// Assign the phone number to the contact and exit the loop
		contact.Phone = phone
		break
	}

	// Return the contact with all details filled in
	return contact
}
