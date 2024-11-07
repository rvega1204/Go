package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// ListContacts retrieves and displays all contacts from the database.
func ListContacts(db *sql.DB) {
	// Define the SQL query to select all records from the 'contact' table.
	query := "SELECT * FROM contact"

	// Execute the query and retrieve the rows from the database.
	rows, err := db.Query(query)
	if err != nil {
		// If there is an error in executing the query, log the error and terminate.
		log.Fatal(err)
	}

	// Ensure that the rows are closed when the function exits.
	defer rows.Close()

	// Print the header for the contact list display.
	fmt.Println("\n Contact List")
	fmt.Println("-----------------------------------------------------------------------")

	// Iterate over the result set, one row at a time.
	for rows.Next() {
		// Create a new contact object to store the current row's data.
		contact := models.Contact{}

		// Declare a variable to store the email value as sql.NullString
		// This allows handling of NULL values in the database email field.
		var emailValue sql.NullString

		// Scan the current row into the contact object and the email value.
		err := rows.Scan(&contact.Id, &contact.Name, &emailValue, &contact.Phone)
		if err != nil {
			// If an error occurs while scanning the row, log the error and terminate.
			log.Fatal(err)
		}

		// Check if the email value is valid (not NULL in the database).
		// If it's valid, assign the email from emailValue to contact.Email.
		// If it's NULL, assign a default value ("No email").
		if emailValue.Valid {
			contact.Email = emailValue.String
		} else {
			contact.Email = "No email"
		}

		// Print the contact details in a formatted manner.
		// Display the contact's ID, name, email, and phone number.
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)

		// Print a separator line between contacts for better readability.
		fmt.Println("-----------------------------------------------------------------------")
	}
}

// GetContactByID retrieves a contact from the database by its ID and displays the contact details.
func GetContactByID(db *sql.DB, contactID int) {
	// Define the SQL query to retrieve the contact by its ID.
	// The "?" placeholder will be replaced by the provided contactID.
	query := "SELECT * FROM contact WHERE id = ?"

	// Execute the query with the contactID as the parameter.
	// QueryRow is used because we expect a single result (one row or none).
	row := db.QueryRow(query, contactID)

	// Initialize an empty contact object to hold the result.
	contact := models.Contact{}

	// Declare a variable to store the email value as sql.NullString.
	// sql.NullString allows us to handle the case where the email might be NULL in the database.
	var emailValue sql.NullString

	// Scan the row into the contact object and the email value.
	// This maps the database columns into the contact fields.
	err := row.Scan(&contact.Id, &contact.Name, &emailValue, &contact.Phone)
	if err != nil {
		// If an error occurs, we handle it:
		// If the error is sql.ErrNoRows, it means no contact with the given ID was found.
		// We log the error and terminate the program with a message.
		if err == sql.ErrNoRows {
			log.Fatalf("No contact found with ID: %d", contactID)
		}
		// If the error is something else, we should handle it properly (optional).
		log.Fatal(err)
	}

	// Check if the email value is valid (not NULL in the database).
	// If valid, assign the email to the contact object.
	// If the email is NULL in the database, set the email to a default value ("No email").
	if emailValue.Valid {
		contact.Email = emailValue.String
	} else {
		contact.Email = "No email"
	}

	// Print the contact details to the console.
	// This includes a header and formatted output for the contact's ID, name, email, and phone number.
	fmt.Println("\n Contact")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)

	// Print a separator line for better visual separation of contact records.
	fmt.Println("-----------------------------------------------------------------------")
}

// CreateContact adds a new contact to the 'contact' table in the database.
func CreateContact(db *sql.DB, contact models.Contact) {
	// Define the SQL query to insert a new contact into the 'contact' table.
	// The query uses placeholders (?) to safely insert values for name, email, and phone.
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	// Execute the query with the contact details passed as arguments.
	// The values are safely inserted into the query using db.Exec.
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)

	// If there is an error executing the query, log the error and terminate the program.
	if err != nil {
		log.Fatal(err)
	}

	// Log a success message indicating that the contact has been successfully added.
	log.Println("New contact added")
}

// UpdateContact updates an existing contact in the 'contact' table based on the contact ID.
func UpdateContact(db *sql.DB, contact models.Contact) {
	// Define the SQL query to update an existing contact.
	// The query sets the new values for 'name', 'email', and 'phone' where the 'id' matches the contact's ID.
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	// Execute the query, passing the updated contact details and the contact ID.
	// The values are safely inserted into the query using placeholders (?).
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)

	// If there is an error executing the query, log the error and terminate the program.
	if err != nil {
		log.Fatal(err)
	}

	// Log a success message indicating that the contact has been successfully updated.
	log.Println("Contact updated")
}

// DeleteContact deletes a contact from the 'contact' table by its ID.
func DeleteContact(db *sql.DB, contactID int) {
	// Define the SQL query to delete a contact based on the provided contact ID.
	// The query uses the placeholder (?) to safely insert the contactID into the query.
	query := "DELETE FROM contact WHERE id = ?"

	// Execute the query, passing the contactID as the parameter to the placeholder (?).
	// This will delete the contact with the matching ID from the database.
	_, err := db.Exec(query, contactID)

	// If an error occurs during query execution, log the error and terminate the program.
	// The program will stop and output the error message.
	if err != nil {
		log.Fatal(err)
	}

	// Log a success message indicating that the contact has been successfully deleted.
	log.Println("Contact deleted")
}
