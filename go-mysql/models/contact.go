package models

// Contact represents a contact in the system with basic information.
// This struct is used to hold the contact details retrieved from or to be inserted into the database.
type Contact struct {
	// Id represents the unique identifier for each contact.
	// It is usually auto-generated by the database when the contact is created.
	Id int

	// Name represents the full name of the contact.
	// This is typically a string of the contact's first and last name.
	Name string

	// Email represents the contact's email address.
	// It is stored as a string and could be used to contact the person electronically.
	Email string

	// Phone represents the contact's phone number.
	// This can be a mobile or landline number and is stored as a string to accommodate various phone number formats.
	Phone string
}
