package main

import (
	"library/animal"
)

func main() {
	//var myBook = book.NewBook("Moby Dick", "Herman Melville", 300)
	// var myBook = book.Book{
	// 	Title:  "Moby Dick",
	// 	Author: "Herman Melville",
	// 	Pages:  300,
	// }

	// myBook.PrintInfo()

	// myBook.SetTitle("Moby Dick (Edicion Especial)")
	// fmt.Println(myBook.GetTitle())

	// var myTextBook = book.NewTextBook("TextBook", "Writter", 300, "Santillana", "Secundaria")

	// book.Print(myBook)
	// book.Print(myTextBook)

	// miPerro := animal.Perro{Nombre: "perrin"}
	// miGato := animal.Gato{Nombre: "gatin"}

	// animal.HacerSonido(&miPerro)
	// animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
