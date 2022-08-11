package main

import "fmt"

func factorialLoop(number int) int {
	var total int = 1
	for i := number; i > 0; i-- {
		total *= i
	}
	return total
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {
	/*
		-- Recursive Function
		1. Recursive function adalah function yang memanggil function diri sendiri.
		2. Recursive function bisa memanggil function diri sendiri lebih dari satu kali.
		3. Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial

		Note: Recursive function harus memiliki kondisi berhenti. Jika tidak, maka akan menghasilkan error/infinite loop.
	*/
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(5))
}
