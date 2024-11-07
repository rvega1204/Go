package book

import "fmt"

// Printable is an interface that defines the method PrintInfo().
// Any type that implements this method satisfies the Printable interface.
type Printable interface {
	PrintInfo() // Method to print details of the object
}

// Print is a function that accepts a Printable type and calls its PrintInfo method.
// This demonstrates polymorphism: any type that implements the Printable interface can be passed.
func Print(p Printable) {
	p.PrintInfo() // Calls the PrintInfo method on the Printable object
}

// Book represents a book with a title, author, and number of pages.
// Fields are private (start with lowercase letter) to encourage encapsulation.
type Book struct {
	title  string // Title of the book
	author string // Author of the book
	pages  int    // Number of pages in the book
}

// NewBook is a constructor function that creates and returns a pointer to a new Book instance.
// It takes the title, author, and number of pages as parameters.
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// SetTitle sets the title of the book.
// This method allows us to modify the title after the Book has been created.
func (b *Book) SetTitle(title string) {
	b.title = title // Sets the book title to the provided value
}

// GetTitle returns the title of the book.
// This method allows us to access the title of the book.
func (b *Book) GetTitle() string {
	return b.title // Returns the book's title
}

// PrintInfo is a method that prints the details of the Book, including title, author, and pages.
// This satisfies the Printable interface and allows the Book to be printed using the Print function.
func (b *Book) PrintInfo() {
	// Prints the book details to the standard output
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// TextBook is a struct that represents a specialized type of Book.
// It includes additional fields such as the editorial and level (e.g., high school, college).
// This struct embeds the Book struct to inherit its properties.
type TextBook struct {
	Book             // Embedding the Book struct, allowing TextBook to inherit Book's fields and methods
	editorial string // The editorial (publisher) of the textbook
	level     string // The academic level for the textbook (e.g., elementary, high school)
}

// NewTextBook is a constructor function that creates and returns a pointer to a new TextBook instance.
// It takes the title, author, pages, editorial, and level as parameters.
func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages}, // Initialize the embedded Book struct
		editorial: editorial,                  // Set the editorial for the textbook
		level:     level,                      // Set the academic level for the textbook
	}
}

// PrintInfo is an overridden method for the TextBook struct that provides more specific details.
// This method is an implementation of the PrintInfo method from the Printable interface for the TextBook type.
func (b *TextBook) PrintInfo() {
	// Prints the detailed information of the textbook, including inherited and new fields
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n", b.title, b.author, b.pages,
		b.editorial, b.level)
}
