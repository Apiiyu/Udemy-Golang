package main

import "fmt"

type HasName interface {
	GetName() string // --> Interface method
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string { // --> This struct method will be used as interface method
	return person.Name
}

type Animal struct {
	Name string
	Type string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	/*
		-- Interface
		1. Interface adalah sebuah tipe data Abstract, dia tidak memiliki implementasi langsung.
		2. Sebuah interface berisikan definisi-definisi method.
		3. Biasanya interface digunakan sebagai kontrak.
		4. Interface juga bisa digunakan didalam parameter.

			-- Implenetasi Interface
			1. Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut.
	*/
	person := Person{
		Name: "John Doe",
	}

	animal := Animal{
		Name: "Kucing",
		Type: "Kucing",
	}
	SayHello(animal)
	SayHello(person)
}
