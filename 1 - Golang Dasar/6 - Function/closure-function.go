package main

import "fmt"

func main() {
	/*
		-- Closure Function
		1. Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama.

		Note: Harap gunakan fitur ini dengan bijak.
	*/

	counter := 0
	secCounter := 0
	increment := func() { //--> Closure Function
		counter++ //--> This will be use vairable counter in outer scope. because in this scope counter is not defined.
		secCounter := 1
		secCounter++
	}

	increment()
	fmt.Println("Counter:", counter)
	fmt.Println("SecCounter:", secCounter)
}
