package animal

import "fmt"

// Define an interface called Animal with a method Sonido
// The Sonido method is expected to be implemented by any type that satisfies the Animal interface.
type Animal interface {
	Sonido() // Method signature
}

// Define a struct called Perro (Dog) which will represent a dog.
type Perro struct {
	Nombre string // Name of the dog
}

// Implement the Sonido method for Perro (Dog).
// This method prints the dog's name followed by " hace guau" (barks).
func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau") // "Nombre hace guau" (Dog's name barks)
}

// Define a struct called Gato (Cat) which will represent a cat.
type Gato struct {
	Nombre string // Name of the cat
}

// Implement the Sonido method for Gato (Cat).
// This method prints the cat's name followed by " hace miau" (meows).
func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau") // "Nombre hace miau" (Cat's name meows)
}

// HacerSonido is a function that takes an Animal as an argument.
// It invokes the Sonido method on the given Animal, making the animal produce its sound.
func HacerSonido(animal Animal) {
	animal.Sonido() // Call the Sonido method on the animal passed to this function
}
